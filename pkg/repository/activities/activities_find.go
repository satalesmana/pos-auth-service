package activities_repo

import (
	"context"
	"go-grpc-auth-svc/pkg/models"
)

func (r *Repo) Find(ctx context.Context, req models.Activities) ([]models.Activities, error) {
	return nil, nil
}
