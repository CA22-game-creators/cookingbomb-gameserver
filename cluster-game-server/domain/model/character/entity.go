package domain

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

type Character struct {
	data *pb.Character
}

func FromRepository(c *pb.Character) Character {
	return Character{
		data: c,
	}
}
