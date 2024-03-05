package users

import (
	"context"
	"go-grpc-auth-svc/pkg/models"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Handler {
	return &Repo{
		db: db,
	}
}

type Handler interface {
	GetByEmail(ctx context.Context, email string) (models.User, error)
	FilterUser(ctx context.Context, req models.User) ([]models.User, error)
	CreateUser(ctx context.Context, req models.User) error
	UpdateUser(ctx context.Context, req models.User) (int, error)
	DeleteUser(ctx context.Context, req models.User) (int, error)
}
