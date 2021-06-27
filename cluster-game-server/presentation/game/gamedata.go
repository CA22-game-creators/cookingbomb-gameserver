package game

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

func (g Service) GameDataStream(stream pb.GameServices_GameDataStreamServer) error {
	return nil
}
