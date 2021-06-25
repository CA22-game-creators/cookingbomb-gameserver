package game

import (
	"context"

	errors "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	auth "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/auth"
	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/session"
	validator "github.com/CA22-game-creators/cookingbomb-proto/server/validation"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

func (g *GameService) Connect(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {

	//TODO: VALIDATE
	if err := validator.Validate(in); err != nil {
		return nil, errors.InvalidArgument()
	}

	token := in.GetSessionToken()
	success, err := auth.AuthToken(token)
	if err != nil {
		return nil, err
	}

	if !success {
		return nil, errors.Unauthorized()
	}

	err = session.ActivateSession(token)
	if err != nil {
		return nil, err
	}

	return &pb.ConnectionResponse{
		Status: session.GetSessionStatus(token),
	}, nil
}

func (g *GameService) Disconnect(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {

	if err := validator.Validate(in); err != nil {
		return nil, errors.InvalidArgument()
	}

	token := in.GetSessionToken()
	allow := auth.CheckToken(token)
	if !allow {
		return nil, errors.SessionNotActive()
	}

	err := session.EndSessionByClient(token)

	if err != nil {
		return nil, err
	}

	return &pb.ConnectionResponse{
		Status: session.GetSessionStatus(token),
	}, nil
}

func (g *GameService) GetConnectionStatus(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {

	if err := validator.Validate(in); err != nil {
		return nil, errors.InvalidArgument()
	}

	token := in.GetSessionToken()

	stats := session.GetSessionStatus(token)

	if stats == pb.ConnectionStatusEnum_CONNECTION_UNSPECIFIED {
		return nil, errors.NoStatusFound()
	}

	return &pb.ConnectionResponse{
		Status: stats,
	}, nil
}
