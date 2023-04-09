package main

import (
	"context"
	"github.com/Portfolio-Adv-Software/PoCs/PoC_gRPC/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	message := proto.Message{
		Body: "Hello from client!",
	}
	c := proto.NewChatServiceClient(conn)
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)
}
