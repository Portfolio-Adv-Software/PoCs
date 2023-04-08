package main

import (
	"github.com/Portfolio-Adv-Software/PoCs/PoC_gRPC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := proto.UnimplementedChatServiceServer{}
	grpcServer := grpc.NewServer()
	proto.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
