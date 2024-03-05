package activities

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
	GetActivities(ctx context.Context) (*pb.ActivitiesResponse, error)
	GetUserActivities(ctx context.Context, req *pb.GetUserActivitiesRequest) (*pb.ActivitiesResponse, error)
}
