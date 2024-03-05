package permissions_repo

import (
	"context"
	"go-grpc-auth-svc/pkg/models"
)

func (r *Repo) UpdateBatch(ctx context.Context, req []models.Permissions, userId int64) error {
	return nil
}
