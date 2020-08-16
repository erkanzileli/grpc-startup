package sum

import (
	"context"
	"grpc-startup/sum/sumpb"
	"log"
)

type Service interface {
	Sum(ctx context.Context, request *sumpb.SumRequest) (*sumpb.SumResponse, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) Sum(ctx context.Context, request *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	result := request.Number1 + request.Number2
	log.Println("Summing", request.Number1, "with", request.Number2, "equals to", result)
	return &sumpb.SumResponse{Result: result}, nil
}
