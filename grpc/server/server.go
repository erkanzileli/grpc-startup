package server

import (
	"grpc-startup/sum"
	"grpc-startup/sum/sumpb"
	"context"
)

type server struct {
	sumService sum.Service
}

func NewServer(sumService sum.Service) *server {
	return &server{sumService: sumService}
}

func (s server) Sum(ctx context.Context, request *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	return s.sumService.Sum(ctx, request)
}
