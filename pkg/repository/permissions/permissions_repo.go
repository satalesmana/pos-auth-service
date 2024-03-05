package permissions_repo

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
	Find(ctx context.Context, req models.Permissions) ([]models.Permissions, error)
	UpdateBatch(ctx context.Context, req []models.Permissions, userId int64) error
	Delete(ctx context.Context, req models.Activities) error
}
