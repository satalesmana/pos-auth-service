package permissions

import (
	"context"
	pb "go-grpc-auth-svc/files/proto/gen"
	"go-grpc-auth-svc/pkg/repository/users"

	"gorm.io/gorm"
)

type Uscase struct {
	userRepo users.Handler
}

func NewAuthUcHandler(conn *gorm.DB) Handler {
	return &Uscase{
		userRepo: users.NewUserRepository(conn),
	}
}

type Handler interface {
	SetUserPermissions(ctx context.Context, req *pb.SetUserPermissionsRequest) (*pb.SetUserPermissionsResponse, error)
}
