package activities_repo

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
	Find(ctx context.Context, req models.Activities) ([]models.Activities, error)
	Create(ctx context.Context, req models.Activities) error
	Update(ctx context.Context, req models.Activities) error
	Delete(ctx context.Context, req models.Activities) error
}
