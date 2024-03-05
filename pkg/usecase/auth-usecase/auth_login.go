package authusecase

import (
	"context"
	pb "go-grpc-auth-svc/files/proto/gen"
	"go-grpc-auth-svc/helpers/hash"
	"go-grpc-auth-svc/helpers/jwt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (uc *Uscase) Login(ctx context.Context, request *pb.LoginRequest, jwt jwt.JwtWrapper) (*pb.LoginResponse, error) {

	user, _ := uc.userRepo.GetByEmail(ctx, request.Email)
	if user.Email == "" {
		err := status.Error(codes.NotFound, "Invalid username or password")
		return nil, err
	}

	match := hash.CheckPasswordHash(request.Password, user.Password)
	if !match {
		err := status.Error(codes.NotFound, "Password doesn't match")
		return nil, err
	}

	token, _ := jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Token: token,
	}, nil
}
