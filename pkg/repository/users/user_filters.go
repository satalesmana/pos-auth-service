package users

import (
	"context"
	"go-grpc-auth-svc/pkg/models"
)

func (r *Repo) FilterUser(ctx context.Context, req models.User) ([]models.User, error) {
	var user []models.User

	r.db.Model(models.User{Email: req.Email}).Find(&user)

	return user, nil
}
