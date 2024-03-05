package users

import (
	"context"
	"go-grpc-auth-svc/helpers/hash"
	"go-grpc-auth-svc/pkg/models"
)

func (r *Repo) CreateUser(ctx context.Context, req models.User) error {
	var user models.User
	var err error

	user.Email = req.Email
	user.Password = hash.HashPassword(req.Password)

	tx := r.db.Begin()
	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
