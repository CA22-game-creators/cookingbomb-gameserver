package game

import (
	"context"

	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/auth"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/session"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	validator "github.com/CA22-game-creators/cookingbomb-proto/server/validation"
)

func (g *Service) Connect(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {

	//TODO: VALIDATE
	if err := validator.Validate(in); err != nil {
		return nil, errors.InvalidArgument(err)
	}

	token := in.GetSessionToken()
	success, err := auth.CheckTokenPermission(token)
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

func (g *Service) Disconnect(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {

	if err := validator.Validate(in); err != nil {
		return nil, errors.InvalidArgument(err)
	}

	token := in.GetSessionToken()
	allow := session.CheckSessionActive(token)
	if !allow {
		return nil, errors.SessionNotActive()
	}

	if err := session.EndSessionByClient(token); err != nil {
		return nil, err
	}

	return &pb.ConnectionResponse{
		Status: session.GetSessionStatus(token),
	}, nil
}

func (g *Service) GetConnectionStatus(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {

	if err := validator.Validate(in); err != nil {
		return nil, errors.InvalidArgument(err)
	}

	token := in.GetSessionToken()

	status := session.GetSessionStatus(token)

	if status == pb.ConnectionStatusEnum_CONNECTION_UNSPECIFIED {
		return nil, errors.NoStatusFound()
	}

	return &pb.ConnectionResponse{
		Status: status,
	}, nil
}
