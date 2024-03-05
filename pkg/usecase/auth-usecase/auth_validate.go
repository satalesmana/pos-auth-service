package authusecase

import (
	"context"
	pb "go-grpc-auth-svc/files/proto/gen"
	"go-grpc-auth-svc/helpers/jwt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (uc Uscase) Validate(ctx context.Context, request *pb.ValidateRequest, jwt jwt.JwtWrapper) (*pb.ValidateResponse, error) {
	claims, err := jwt.ValidateToken(request.Token)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	user, err := uc.userRepo.GetByEmail(ctx, claims.Email)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.ValidateResponse{
		UserId: user.Id,
		Email:  user.Email,
	}, nil
}
