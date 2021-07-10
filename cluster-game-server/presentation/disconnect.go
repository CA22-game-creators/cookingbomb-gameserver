package presentation

import (
	"context"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	validator "github.com/CA22-game-creators/cookingbomb-proto/server/validation"

	application "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/disconnect"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
)

func (c controller) Disconnect(ctx context.Context, req *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, errors.InvalidArgument(err.Error())
	}

	input := application.InputData{SessionToken: req.GetSessionToken()}
	output := c.disconnect.Handle(input)
	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.ConnectionResponse{
		Status: StatusMapper(output.Status),
	}, nil
}
