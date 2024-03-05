package services

import (
	"context"

	pb "go-grpc-auth-svc/files/proto/gen"
	"go-grpc-auth-svc/helpers/jwt"
	uc "go-grpc-auth-svc/pkg/usecase/auth-usecase"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type Server struct {
	ConDb *gorm.DB
	Jwt   jwt.JwtWrapper
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	usersUc := uc.NewAuthUcHandler(s.ConDb)
	result := usersUc.Register(context.Background(), req)

	return result, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	usersUc := uc.NewAuthUcHandler(s.ConDb)
	result := usersUc.Login(context.Background(), req, s.Jwt)

	return result, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	usersUc := uc.NewAuthUcHandler(s.ConDb)
	result := usersUc.Validate(context.Background(), req, s.Jwt)

	return result, nil
}

func (s *Server) SetUserPermissions(ctx context.Context, req *pb.SetUserPermissionsRequest) (*pb.SetUserPermissionsResponse, error) {
	return nil, nil
}

func (s *Server) GetActivities(ctx context.Context, req *emptypb.Empty) (*pb.ActivitiesResponse, error) {
	return nil, nil
}

func (s *Server) GetUserActivities(ctx context.Context, req *pb.GetUserActivitiesRequest) (*pb.ActivitiesResponse, error) {
	return nil, nil
}
