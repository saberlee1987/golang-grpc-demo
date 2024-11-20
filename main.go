package main

import (
	"context"
	"fmt"
	"golang-grpc-demo/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloService struct {
	proto.HelloServiceServer
}

func (h *HelloService) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("request sayHello firstName : %s , lastName : %s\n", request.GetFirstName(), request.GetLastName())
	message := fmt.Sprintf("Hello %s %s", request.GetFirstName(), request.GetLastName())
	log.Printf("request sayHello firstName : %s , lastName : %s , response ===> %s\n",
		request.GetFirstName(), request.GetLastName(), message)
	return &proto.HelloResponse{Message: message}, nil
}

func main() {
	tcp, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcServer, &HelloService{})
	err = grpcServer.Serve(tcp)
	if err != nil {
		log.Fatal(err)
	}
}
