package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"project-go-gRPC/gRPC"
	"project-go-gRPC/internal/service"
)

func main() {
	serv := &service.SumNumber{}
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
