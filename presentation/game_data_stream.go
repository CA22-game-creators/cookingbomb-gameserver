package presentation

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"

	application "github.com/CA22-game-creators/cookingbomb-gameserver/application/game_data_stream"
)

func (c controller) GameDataStream(stream pb.GameServices_GameDataStreamServer) error {
	input := application.InputData{Stream: stream}
	output := c.gameDataStream.Handle(input)
	return output.Err
}
