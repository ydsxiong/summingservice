package impl

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/ydsxiong/summingservice/gRPC/domain"
	"github.com/ydsxiong/summingservice/gRPC/service"
)

type serverMiddleware func(server service.SumServiceServer) service.SumServiceServer

type loggingMiddlewareService struct {
	logger log.Logger
	svc    service.SumServiceServer
}

func NewLoggingMiddlewareService(logger log.Logger) serverMiddleware {
	return func(service service.SumServiceServer) service.SumServiceServer {
		return loggingMiddlewareService{logger, service}
	}
}

func (lm loggingMiddlewareService) Sum(ctx context.Context, req *domain.SumRequest) (output *domain.SumResponse, err error) {
	defer func(begin time.Time) {
		_ = lm.logger.Log(
			"service_call", "sum",
			"Input_ID", req.Name, "Numbers", req.Numbers,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return lm.svc.Sum(ctx, req)
}

func (lm loggingMiddlewareService) GetAllSums(req *domain.SumFilter, srv service.SumService_GetAllSumsServer) (err error) {
	defer func(begin time.Time) {
		_ = lm.logger.Log(
			"service_call", "getAllSums",
			"Filter_by_keyword", req.Keyword,
			"output", "stream of messages",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return lm.svc.GetAllSums(req, srv)
}

type instrumentingMiddlewareService struct {
	ints  metrics.Counter
	chars metrics.Counter
	svc   service.SumServiceServer
}

func NewInstrumentingMiddlewareService(ints, chars metrics.Counter) serverMiddleware {
	return func(service service.SumServiceServer) service.SumServiceServer {
		return instrumentingMiddlewareService{ints, chars, service}
	}
}

func (ism instrumentingMiddlewareService) Sum(ctx context.Context, req *domain.SumRequest) (output *domain.SumResponse, err error) {
	output, err = ism.svc.Sum(ctx, req)
	if output != nil {
		//ism.ints.Add(float64(output.Sum))
	}
	return
}

func (ism instrumentingMiddlewareService) GetAllSums(req *domain.SumFilter, srv service.SumService_GetAllSumsServer) (err error) {
	err = ism.svc.GetAllSums(req, srv)
	if err == nil {
		//ism.chars.Add(float64(srv.NumOfMessagesSent()))
	}
	return
}
