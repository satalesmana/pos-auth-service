package permissions_repo

import (
	"context"
	"go-grpc-auth-svc/pkg/models"
)

func (r *Repo) Find(ctx context.Context, req models.Permissions) ([]models.Permissions, error) {
	return nil, nil
}
