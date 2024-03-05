package authusecase

import (
	"context"
	pb "go-grpc-auth-svc/files/proto/gen"
	"go-grpc-auth-svc/pkg/models"

	"google.golang.org/grpc/codes"
)

func (uc *Uscase) Register(ctx context.Context, request *pb.RegisterRequest) *pb.RegisterResponse {
	var userRequest models.User
	userRequest.Email = request.Email
	userRequest.Password = request.Password

	user, _ := uc.userRepo.GetByEmail(ctx, request.Email)
	if user.Email != "" {
		return &pb.RegisterResponse{
			Status: int64(codes.AlreadyExists),
			Error:  "E-Mail already exists",
		}
	}

	uc.userRepo.CreateUser(ctx, userRequest)

	return &pb.RegisterResponse{
		Status: int64(codes.OK),
	}
}
