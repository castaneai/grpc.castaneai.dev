package main

import (
	"context"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/castaneai/grpc.castaneai.dev/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func usage() {
	log.Printf("Usage: ./testclient <host:port> <num_stream>")
	os.Exit(2)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	addr := os.Args[1]
	numStream, err := strconv.Atoi(os.Args[2])

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
	for i := 1; i <= numStream; i++ {
		go startStream(ctx, echo, fmt.Sprintf("stream%03d", i))
	}
	<-ctx.Done()
}

func startStream(ctx context.Context, c proto.EchoServiceClient, streamID string) {
	stream, err := c.StreamingEcho(ctx)
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
	ticker := time.NewTicker(2 * time.Second)
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			msg := fmt.Sprintf("%s-%d", streamID, i)
			if err := stream.Send(&proto.StreamingEchoRequest{Message: msg}); err != nil {
				log.Fatalf("failed to send message: %+v", err)
			}
			log.Printf("sent: %s", msg)
			i++
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
