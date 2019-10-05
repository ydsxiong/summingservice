package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ydsxiong/summingservice/cmd/gRPC/client/cli"
	"github.com/ydsxiong/summingservice/gRPC/service"
	"google.golang.org/grpc"
)

const (
	GRPC_SERVER_HOST = "GRPC_SERVER_HOST"
)

func getServerHostFromEnvironment() (hostAddr string, err error) {
	hostAddr, provided := os.LookupEnv(GRPC_SERVER_HOST)
	if !provided || hostAddr == "" {
		// var host = flag.String("grpc_server_host", "localhost:9000", "remote grpc server address")
		// flag.Parse()
		// grpc_server_host = *host
		err = fmt.Errorf("Missing grpc server host from environment")
	}
	return
}

/**
allow user to play with the sum service on cmd line:
*/
func main() {
	hostAddr, err := getServerHostFromEnvironment()
	if err != nil {
		log.Fatalln(err)
	}
	conn, e := grpc.Dial(hostAddr, grpc.WithInsecure())

	if e != nil {
		log.Fatalln(e)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	gRPCClient := service.NewSumServiceClient(conn)

	cli.NewSumServiceCLI(os.Stdin, os.Stdout, gRPCClient, ctx).Run()
}
