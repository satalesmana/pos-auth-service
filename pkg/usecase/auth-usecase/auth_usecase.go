package authusecase

import (
	"context"
	pb "go-grpc-auth-svc/files/proto/gen"
	"go-grpc-auth-svc/helpers/jwt"
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
	Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, request *pb.LoginRequest, jwt jwt.JwtWrapper) (*pb.LoginResponse, error)
	Validate(ctx context.Context, request *pb.ValidateRequest, jwt jwt.JwtWrapper) (*pb.ValidateResponse, error)
}
