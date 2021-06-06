package game

import (
	"context"

	auth "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

func (g *GameService) Connect(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	token := in.GetSessionToken()
	success, err := auth.AuthToken(token)
	if err != nil {
		return nil, err
	}
	if !success {
		return &pb.ConnectionResponse{
			Status: pb.ConnectionStatusEnum_CONNECTION_FAIL,
		}, status.Errorf(codes.PermissionDenied, "Token Rejected")
	}

	//TODO: token status check

	return &pb.ConnectionResponse{
		Status: pb.ConnectionStatusEnum_CONNECTED,
	}, nil
}

func (g *GameService) Disconnect(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	token := in.GetSessionToken()
	allow, err := auth.CheckToken(token)
	if err != nil {
		return nil, err
	}
	if !allow {
		return nil, status.Errorf(codes.PermissionDenied, "Token Rejected")
	}

	//TODO: token status update

	return &pb.ConnectionResponse{
		Status: pb.ConnectionStatusEnum_DISCONNECTED_BY_CLIENT,
	}, nil
}

func (g *GameService) GetConnectionStatus(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	token := in.GetSessionToken()
	allow, err := auth.CheckToken(token)
	if err != nil {
		return nil, err
	}
	if !allow {
		return nil, status.Errorf(codes.PermissionDenied, "Token Rejected")
	}

	//TODO: token status check

	return &pb.ConnectionResponse{
		Status: pb.ConnectionStatusEnum_CONNECTED,
	}, nil
}
