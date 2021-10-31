package application

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

type InputData struct {
	Stream pb.GameServices_GameDataStreamServer
}
