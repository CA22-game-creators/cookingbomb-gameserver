package account

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
)

type Account struct {
	id   string
	name string
}

func NewFromAPIResponce(res *pb.AccountInfo) Account {
	return Account{
		id:   res.GetId(),
		name: res.GetName(),
	}
}
