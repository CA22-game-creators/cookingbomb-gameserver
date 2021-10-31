package presentation

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"

	gameDataStream "github.com/CA22-game-creators/cookingbomb-gameserver/application/game_data_stream"
)

type controller struct {
	gameDataStream gameDataStream.InputPort
}

func New(g gameDataStream.InputPort) pb.GameServicesServer {
	return controller{
		gameDataStream: g,
	}
}
