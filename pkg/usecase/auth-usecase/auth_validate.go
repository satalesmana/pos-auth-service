package authusecase

import (
	"context"
	pb "go-grpc-auth-svc/files/proto/gen"
	"go-grpc-auth-svc/helpers/jwt"

	"google.golang.org/grpc/codes"
)

func (uc Uscase) Validate(ctx context.Context, request *pb.ValidateRequest, jwt jwt.JwtWrapper) *pb.ValidateResponse {
	claims, err := jwt.ValidateToken(request.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: int64(codes.InvalidArgument),
			Error:  err.Error(),
		}
	}

	user, err := uc.userRepo.GetByEmail(ctx, claims.Email)
	if err != nil {
		return &pb.ValidateResponse{
			Status: int64(codes.InvalidArgument),
			Error:  err.Error(),
		}
	}

	return &pb.ValidateResponse{
		Status: int64(codes.OK),
		Error:  "",
		UserId: user.Id,
		Email:  user.Email,
	}
}
