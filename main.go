package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc.castaneai.dev/proto"
)

func main() {
	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = fmt.Sprintf(":%s", port)
	}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}
	s := grpc.NewServer()
	proto.RegisterEchoServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("listening on %s...", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}

type server struct {
	proto.UnimplementedEchoServiceServer
}

func (s server) StreamingEcho(stream proto.EchoService_StreamingEchoServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := stream.Send(&proto.StreamingEchoResponse{
			Message: req.Message,
		}); err != nil {
			return err
		}
	}

}
