package application

import (
	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
)

type interactor struct {
	repository domain.Repository
}

func New(r domain.Repository) InputPort {
	return &interactor{
		repository: r,
	}
}

func (i interactor) Handle(input InputData) OutputData {

	val, err := i.repository.Find(input.SessionToken)
	if err != nil {
		return OutputData{}
	}

	userid := val.ID
	return OutputData{
		Status: i.repository.GetSessionStatus(userid),
	}
}
