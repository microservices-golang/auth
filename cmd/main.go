package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/microservices-golang/auth/pkg/user"
)

const grpcPort = 50051

type server struct {
	user.UnimplementedUserServiceServer
}

// GetUser - получить данные о пользователе
func (s *server) GetUser(_ context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
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

// CreateUser - создает нового пользователя
func (s *server) CreateUser(_ context.Context, _ *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return &user.CreateUserResponse{
		Id: 2,
	}, nil
}

// DeleteUser - удаление пользователя
func (s *server) DeleteUser(_ context.Context, req *user.DeleteUserRequest) (*emptypb.Empty, error) {
	fmt.Println(req.GetId())
	return &emptypb.Empty{}, nil
}

// UpdateUser - обновление данных пользователя
func (s *server) UpdateUser(_ context.Context, req *user.UpdateUserRequest) (*emptypb.Empty, error) {
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
