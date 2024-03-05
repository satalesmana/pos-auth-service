package permissions

import (
	"context"
	pb "go-grpc-auth-svc/files/proto/gen"
)

func (uc *Uscase) SetUserPermissions(ctx context.Context, req *pb.SetUserPermissionsRequest) (*pb.SetUserPermissionsResponse, error) {
	return nil, nil
}
