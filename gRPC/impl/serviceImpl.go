package impl

import (
	"context"
	"strings"

	"github.com/ydsxiong/summingservice/datastore"
	"github.com/ydsxiong/summingservice/gRPC/domain"
	"github.com/ydsxiong/summingservice/gRPC/service"
)

type sumServer struct {
	store datastore.Store
}

func NewSumServiceImpl(store datastore.Store) service.SumServiceServer {
	return &sumServer{store}
}

func (s *sumServer) Sum(ctx context.Context, req *domain.SumRequest) (*domain.SumResponse, error) {
	var sum int64
	for _, num := range req.Numbers {
		sum += num
	}
	s.store.Save(datastore.DataMessage{Name: req.Name, Input: req.Numbers, Output: sum})
	return &domain.SumResponse{Input: req, Sum: sum}, nil
}

func (s *sumServer) GetAllSums(req *domain.SumFilter, srv service.SumService_GetAllSumsServer) error {
	msgs, err := s.store.FetchAll()
	if err != nil {
		return err
	}
	if msgs != nil {
		for _, msg := range msgs {
			if req != nil && !strings.Contains(msg.Name, req.Keyword) {
				continue
			}
			res := new(domain.SumResponse)
			res.Input = &domain.SumRequest{Numbers: msg.Input, Name: msg.Name}
			res.Sum = msg.Output
			if err := srv.Send(res); err != nil {
				return err
			}
		}
	}
	return nil
}
