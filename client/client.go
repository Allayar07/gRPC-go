package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"project-go-gRPC/gRPC"
	"strconv"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatalln("not enough args")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := gRPC.NewSumClient(conn)
	result, err := client.SumNumber(context.Background(), &gRPC.SumRequest{
		X: int32(x),
		Y: int32(y),
	})

	log.Println(result)
}
