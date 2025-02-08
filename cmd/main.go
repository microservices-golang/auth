package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/microservices-golang/auth/pkg/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

//protoc  --go_out=./pkg/user  --go-grpc_out=./pkg/user api/user/auth.proto

type server struct {
	user.UnimplementedUserServiceServer
}

func (s *server) Get(_ context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	id := req.GetId()
	return &user.GetUserResponse{
		Id:        id,
		Name:      "Denis",
		Email:     "sadasda@yandex.ru",
		Role:      user.Role_ADMIN,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Create(_ context.Context, _ *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return &user.CreateUserResponse{
		Id: 2,
	}, nil
}

func (s *server) Delete(_ context.Context, req *user.DeleteUserRequest) (*emptypb.Empty, error) {
	fmt.Println(req.GetId())
	return &emptypb.Empty{}, nil
}

func (s *server) Update(_ context.Context, req *user.UpdateUserRequest) (*emptypb.Empty, error) {
	fmt.Println(req.GetId())
	return &emptypb.Empty{}, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user.RegisterUserServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
