package main

import (
	"context"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/castaneai/grpc.castaneai.dev/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func usage() {
	log.Printf("Usage: ./testclient <host:port>")
	os.Exit(2)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	addr := os.Args[1]

	creds, err := loadCredentials()
	if err != nil {
		log.Fatalf("failed to load client credentials: %+v", err)
	}
	log.Printf("connecting to %s...", addr)
	cc, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect: %+v", err)
	}
	echo := proto.NewEchoServiceClient(cc)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	stream, err := echo.StreamingEcho(ctx)
	if err != nil {
		log.Fatalf("failed to call echo: %+v", err)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				recv, err := stream.Recv()
				if err != nil {
					log.Fatalf("failed to recv message: %+v", err)
				}
				log.Printf("received: %s", recv.Message)
			}
		}
	}()
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			msg := fmt.Sprintf("%v", time.Now())
			if err := stream.Send(&proto.StreamingEchoRequest{Message: msg}); err != nil {
				log.Fatalf("failed to send message: %+v", err)
			}
			log.Printf("sent: %s", msg)
		}
	}
}

func loadCredentials() (credentials.TransportCredentials, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to get system cert pool: %w", err)
	}
	return credentials.NewClientTLSFromCert(certPool, ""), nil
}
