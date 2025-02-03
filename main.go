package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "github.com/abhayishere/lokrr-auth-service/proto/authService"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAuthServiceServer
}

func (s *server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	fmt.Println("User Registered:", req.Email)
	userID := uuid.New().String()
	return &proto.RegisterResponse{
		UserId:  userID,
		Message: "Registration successful",
	}, nil
}

func (s *server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	fmt.Println("User Logged In:", req.Email)
	token := uuid.New().String()
	return &proto.LoginResponse{
		UserId: req.Email,
		Token:  token,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterAuthServiceServer(grpcServer, &server{})

	fmt.Println("Auth Service running on port 50052...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
