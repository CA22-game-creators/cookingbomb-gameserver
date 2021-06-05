package connection

import (
	"context"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

type connection struct{}

func (c connection) Connection(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	return &pb.ConnectionResponse{
		Status: pb.ConnectionStatusEnum_CONNECTED,
	}, nil
}
