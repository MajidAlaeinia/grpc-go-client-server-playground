package chat

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	log.Printf("Received message id from client: %d", message.Id)
	return &Message{Body: "Hello From the Server!"}, nil
}
