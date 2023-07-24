package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"project-go-gRPC/gRPC"
)

type sumServer struct {
	gRPC.UnimplementedSumServer
}

func (s *sumServer) SumNumber(ctx context.Context, req *gRPC.SumRequest) (*gRPC.SumResponse, error) {
	return &gRPC.SumResponse{
		Result: req.GetX() + req.GetY(),
	}, nil
}

func main() {
	serv := &sumServer{}

	gServer := grpc.NewServer()

	gRPC.RegisterSumServer(gServer, serv)

	listiner, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalln(err)
	}
	if err = gServer.Serve(listiner); err != nil {
		log.Fatalln(err)
	}
}
