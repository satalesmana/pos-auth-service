package authusecase

import (
	"context"
	pb "go-grpc-auth-svc/files/proto/gen"
	"go-grpc-auth-svc/helpers/hash"
	"go-grpc-auth-svc/helpers/jwt"

	"google.golang.org/grpc/codes"
)

func (uc *Uscase) Login(ctx context.Context, request *pb.LoginRequest, jwt jwt.JwtWrapper) *pb.LoginResponse {

	user, _ := uc.userRepo.GetByEmail(ctx, request.Email)
	if user.Email == "" {
		return &pb.LoginResponse{
			Status: int64(codes.NotFound),
			Error:  "invalid username or password",
		}
	}

	match := hash.CheckPasswordHash(request.Password, user.Password)
	if !match {
		return &pb.LoginResponse{
			Status: int64(codes.NotFound),
			Error:  "password doesn't match",
		}
	}

	token, _ := jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: int64(codes.OK),
		Error:  "",
		Token:  token,
	}
}
