package user

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

type User struct {
	id     string
	name   string
	status pb.ConnectionStatusEnum
}

func (u *User) GetStatus() pb.ConnectionStatusEnum {
	return u.status
}
