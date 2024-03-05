package users

import (
	"context"
	"go-grpc-auth-svc/pkg/models"
)

func (r *Repo) GetByEmail(ctx context.Context, email string) (models.User, error) {
	var (
		user models.User
		err  error
	)

	tx := r.db.Begin()
	if err = tx.Find(&user, "email = ?", email).Error; err != nil {
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, err
}
