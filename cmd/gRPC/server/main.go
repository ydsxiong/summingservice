package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/ydsxiong/summingservice/datastore"
	"github.com/ydsxiong/summingservice/gRPC/impl"
	service "github.com/ydsxiong/summingservice/gRPC/service"
	"google.golang.org/grpc"
)

const (
	SERVICE_TCP_PORT = "SERVICE_TCP_PORT"
)

func getPortFromEnvironment() (int, error) {
	tcp_port, provided := os.LookupEnv(SERVICE_TCP_PORT)
	if !provided || tcp_port == "" {
		// var portstr = flag.String("port", "9000", "http listen address")
		// flag.Parse()
		// tcp_port = *portstr
	}
	port, err := strconv.Atoi(tcp_port)
	if err != nil {
		err = fmt.Errorf("Incorrect port number specified: %s", err)
	}
	return port, err
}

/**
Configure a grpc server to set it up and running
*/
func main() {
	errChan := make(chan error)

	//
	// config gprc server:
	gRPCServer, err := configGRPCServer()
	if err != nil {
		log.Fatalln(err)
	}

	//
	// stop the server for CTRL+C
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// start up the server
	go func() {
		netListener, portstr, err := getNetListener()
		if err != nil {
			errChan <- err
		} else {
			log.Println("Listening on port:", portstr)
			errChan <- gRPCServer.Serve(netListener)
		}
	}()

	log.Fatalln(<-errChan)

}

func getNetListener() (net.Listener, string, error) {
	port, err := getPortFromEnvironment()
	if err != nil {
		return nil, "", err
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, "", fmt.Errorf("failed to listen: %v", err)
	}

	return lis, strconv.Itoa(port), nil
}

func configGRPCServer() (*grpc.Server, error) {
	gRPCServer := grpc.NewServer()
	service.RegisterSumServiceServer(gRPCServer, setupServerHandler())
	return gRPCServer, nil
}

func setupServerHandler() service.SumServiceServer {
	store := datastore.NewInMemoryDataStore() // can be replaced by a proper data store implementing an external service.
	sumServiceHandler := impl.NewSumServiceImpl(store)

	// add some middlewares around the service:
	{
		logger := kitlog.NewLogfmtLogger(os.Stderr)
		logger = kitlog.With(logger, "listen", 7000, "caller", kitlog.DefaultCaller)
		sumServiceHandler = impl.NewLoggingMiddlewareService(logger)(sumServiceHandler)

		fieldKeys := []string{"method", "error"}
		requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "sum_service",
			Name:        "request_count",
			Help:        "Number of requests received.",
			ConstLabels: stdprometheus.Labels{"label1": "foo", "label2": "bar"},
		}, fieldKeys)
		responseCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "sum_service",
			Name:        "response_count",
			Help:        "Number of messages in the response received.",
			ConstLabels: stdprometheus.Labels{"label1": "foo", "label2": "bar"},
		}, fieldKeys)

		sumServiceHandler = impl.NewInstrumentingMiddlewareService(requestCount, responseCount)(sumServiceHandler)
	}

	return sumServiceHandler
}
