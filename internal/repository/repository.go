package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/microservices-golang/auth/pkg/user"
)

// Repository предоставляет доступ к методам в PostgreSQL
type Repository struct {
	pool *pgxpool.Pool
}

const (
	dbDSN = "host=localhost port=54321 dbname=microservice-auth user=bulgakov password=bulgakov sslmode=disable"
)

// NewRepository создает новый экземпляр Repository с пулом соединений к БД
func NewRepository(ctx context.Context) *Repository {
	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return &Repository{pool: pool}
}

// Insert осуществляет запрос в бд для добавления данных о пользователе
func (r *Repository) Insert(ctx context.Context, req *user.CreateUserRequest) (int64, error) {
	var id int64
	query := "INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id"

	err := r.pool.QueryRow(ctx, query, req.Name, req.Email, req.Password, req.Role).Scan(&id)
	if err != nil {
		log.Printf("Ошибка при добавлении пользователя: %v", err)
		return 0, fmt.Errorf("Ошибка при добавлении пользователя: %v", err)
	}
	return id, nil
}

// Get осущестляет запрос в бд для получения данных о пользователе
func (r *Repository) Get(ctx context.Context, reqID int64) (*user.GetUserResponse, error) {
	var id int64
	var name, email string
	var roleStr string
	var createdAt, updatedAt time.Time
	query := `SELECT id, name, email, role, created_at, updated_at 	FROM users 	WHERE id = $1`

	err := r.pool.QueryRow(ctx, query, reqID).Scan(&id, &name, &email, &roleStr, &createdAt, &updatedAt)
	if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return nil, fmt.Errorf("Ошибка при получении пользователя: %v", err)
	}
	createdAtProto := timestamppb.New(createdAt)
	updatedAtProto := timestamppb.New(updatedAt)

	var role user.Role
	switch roleStr {
	case "ADMIN":
		role = user.Role_ADMIN
	case "USER":
		role = user.Role_USER
	default:
		role = user.Role_UNKNOWN
	}

	return &user.GetUserResponse{
		Id: id, Name: name, Email: email, Role: role, CreatedAt: createdAtProto, UpdatedAt: updatedAtProto,
	}, nil
}

// Update осуществляет запрос в бд для обновления данных о пользователе
func (r *Repository) Update(ctx context.Context, req *user.UpdateUserRequest) error {

	query := `UPDATE users SET name = $1, email = $2, role = $3, updated_at = NOW() WHERE id = $4`

	_, err := r.pool.Exec(ctx, query, req.Name, req.Email, req.Role, req.Id)
	if err != nil {
		log.Printf("Ошибка при обновлении данных: %v", err.Error())
		return fmt.Errorf("Ошибка при обновлении данных: %v", err)
	}

	return nil
}

// Delete осуществляет запрос в бд для удаления пользователя по id
func (r *Repository) Delete(ctx context.Context, req *user.DeleteUserRequest) error {

	query := `DELETE FROM users WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, req.Id)
	if err != nil {
		log.Printf("Ошибка при удалении данных: %v", err.Error())
		return fmt.Errorf("Ошибка при удалении данных: %v", err)
	}

	return nil
}
