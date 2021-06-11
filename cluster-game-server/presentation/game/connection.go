package game

import (
	"context"

	auth "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/auth"
	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/session"
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

	session.ActivateSession(token)

	return &pb.ConnectionResponse{
		Status: session.GetSessionStatus(token),
	}, nil
}

func (g *GameService) Disconnect(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	token := in.GetSessionToken()
	allow := auth.CheckToken(token)
	if !allow {
		return nil, status.Errorf(codes.PermissionDenied, "Token Rejected")
	}

	session.EndSessionByClient(token)

	return &pb.ConnectionResponse{
		Status: session.GetSessionStatus(token),
	}, nil
}

func (g *GameService) GetConnectionStatus(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	token := in.GetSessionToken()
	allow := auth.CheckToken(token)
	if !allow {
		return nil, status.Errorf(codes.PermissionDenied, "Token Rejected")
	}

	return &pb.ConnectionResponse{
		Status: session.GetSessionStatus(token),
	}, nil
}
