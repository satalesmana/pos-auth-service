package authusecase

import (
	"context"
	pb "go-grpc-auth-svc/files/proto/gen"
	"go-grpc-auth-svc/pkg/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (uc *Uscase) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var userRequest models.User
	userRequest.Email = request.Email
	userRequest.Password = request.Password

	user, _ := uc.userRepo.GetByEmail(ctx, request.Email)
	if user.Email != "" {
		err := status.Error(codes.AlreadyExists, "E-Mail already exists")
		return nil, err
	}

	uc.userRepo.CreateUser(ctx, userRequest)

	return &pb.RegisterResponse{
		Message: "Users successfully created",
	}, nil
}
