package account

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
)

type Account struct {
	Id   string
	Name string
}

func NewFromAPIResponce(res *pb.AccountInfo) Account {
	return Account{
		Id:   res.GetId(),
		Name: res.GetName(),
	}
}
