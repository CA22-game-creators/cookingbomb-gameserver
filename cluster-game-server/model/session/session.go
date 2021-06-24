package session

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

type Session struct {
	Status pb.ConnectionStatusEnum
}

func (u *Session) GetStatus() pb.ConnectionStatusEnum {
	return u.Status
}

func (u *Session) IsActive() bool {
	return u.Status == pb.ConnectionStatusEnum_CONNECTED
}
