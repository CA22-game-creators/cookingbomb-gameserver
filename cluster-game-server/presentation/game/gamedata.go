package game

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

func (g GameService) GameDataStream(stream pb.GameServices_GameDataStreamServer) error {
	return nil
}
