package main

import (
	pb "github.com/ReStorePUC/protobucket/payment"
	"github.com/restore/payment/config"
	"github.com/restore/payment/handler"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	config.Init()

	// GRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServer(s, handler.NewPaymentServer(config.NewPaymentConfig()))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
