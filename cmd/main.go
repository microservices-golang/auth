package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/microservices-golang/auth/internal/repository"
	"github.com/microservices-golang/auth/internal/service"
	"github.com/microservices-golang/auth/pkg/user"
)

const grpcPort = 50051

func main() {

	ctx := context.Background()

	db := repository.NewRepository(ctx)

	srv := service.NewService(db)

	s := grpc.NewServer()

	reflection.Register(s)

	user.RegisterUserServiceServer(s, srv)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
