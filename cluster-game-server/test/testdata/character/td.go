package testdata

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

var Character = pb.Character{
	Id:       "TestUser",
	IsActive: true,
	Position: &pb.Position{
		X: 0,
		Y: 0,
		Z: 0,
	},
	Rotation: &pb.Rotation{
		X: 0,
		Y: 0,
		Z: 0,
		W: 0,
	},
	Verocity: &pb.Verocity{
		X: 0,
		Y: 0,
		Z: 0,
	},
}
