package proto

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello from the Server!"}, nil
}
