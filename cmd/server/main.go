package main

import (
	"auth-microservice/config"
	"auth-microservice/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	cfg := config.LoadConfig()

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc.JWTAuthInterceptor))
	grpc.RegisterAuthServiceServer(s, &grpc.AuthServer{})
	reflection.Register(s)

	log.Printf("Starting gRPC server on %s", cfg.GRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
