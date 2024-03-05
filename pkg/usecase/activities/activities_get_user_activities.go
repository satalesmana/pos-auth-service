package activities

import (
	"context"
	pb "go-grpc-auth-svc/files/proto/gen"
)

func (uc *Uscase) GetUserActivities(ctx context.Context, req *pb.GetUserActivitiesRequest) (*pb.ActivitiesResponse, error) {
	return nil, nil
}
