package service

import (
	"context"
	"project-go-gRPC/gRPC"
)

type SumNumber struct {
	gRPC.UnimplementedSumServer
}

func (s *SumNumber) SumNumber(_ context.Context, req *gRPC.SumRequest) (*gRPC.SumResponse, error) {
	return &gRPC.SumResponse{
		Result: req.GetX() + req.GetY(),
	}, nil
}
