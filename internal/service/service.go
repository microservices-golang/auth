package service

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/microservices-golang/auth/pkg/user"
)

// DbRepo определяет контракт для работы с базой данных.
type DbRepo interface {
	Insert(ctx context.Context, req *user.CreateUserRequest) (int64, error)
	Get(ctx context.Context, req int64) (*user.GetUserResponse, error)
	Update(ctx context.Context, req *user.UpdateUserRequest) error
	Delete(ctx context.Context, req *user.DeleteUserRequest) error
}

// Service реализует gRPC-сервер для работы с пользователями
type Service struct {
	user.UnimplementedUserServiceServer
	dbR DbRepo
}

// NewService создает новый экземпляр UserService
func NewService(dbR DbRepo) *Service {
	return &Service{dbR: dbR}
}

// GetUser возвращает информацию о пользователе
func (s *Service) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	userData, _ := s.dbR.Get(ctx, req.Id)

	return &user.GetUserResponse{
		Id:        userData.Id,
		Name:      userData.Name,
		Email:     userData.Email,
		Role:      userData.Role,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}, nil
}

// CreateUser - создает нового пользователя
func (s *Service) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	id, _ := s.dbR.Insert(ctx, req)

	return &user.CreateUserResponse{
		Id: id,
	}, nil
}

// DeleteUser - удаление пользователя
func (s *Service) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*emptypb.Empty, error) {
	err := s.dbR.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// UpdateUser - обновление данных пользователя
func (s *Service) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*emptypb.Empty, error) {

	_, err := s.dbR.Get(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("Пользователь не найден")
	}
	err = s.dbR.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
