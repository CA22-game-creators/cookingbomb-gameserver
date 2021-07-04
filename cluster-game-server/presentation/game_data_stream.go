package presentation

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

func (c controller) GameDataStream(stream pb.GameServices_GameDataStreamServer) error {
	return nil
}
