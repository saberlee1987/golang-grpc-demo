package main

import (
	"context"
	"golang-grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	connection, err := grpc.NewClient("localhost:8090", credentials)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	client := proto.NewHelloServiceClient(connection)

	helloRequest := &proto.HelloRequest{
		FirstName: "saber1366",
		LastName:  "Azizi",
	}
	log.Printf("request sayHello firstName : %s , lastName : %s\n",
		helloRequest.GetFirstName(), helloRequest.GetLastName())
	helloResponse, err := client.SayHello(context.Background(), helloRequest)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("request sayHello firstName : %s , lastName : %s , response ===> %s\n",
		helloRequest.GetFirstName(), helloRequest.GetLastName(), helloResponse.GetMessage())
}
