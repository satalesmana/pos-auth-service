package users

import (
	"context"
	"go-grpc-auth-svc/pkg/models"
)

func (r *Repo) DeleteUser(ctx context.Context, req models.User) (int, error) {

	return 0, nil
}
