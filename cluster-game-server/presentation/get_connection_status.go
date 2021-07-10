package presentation

import (
	"context"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	validator "github.com/CA22-game-creators/cookingbomb-proto/server/validation"

	application "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/get_connection_status"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
)

func (c controller) GetConnectionStatus(ctx context.Context, req *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, errors.InvalidArgument(err.Error())
	}

	input := application.InputData{SessionToken: req.GetSessionToken()}
	output := c.getConnectionStatus.Handle(input)

	return &pb.ConnectionResponse{
		Status: StatusMapper(output.Status),
	}, nil
}
