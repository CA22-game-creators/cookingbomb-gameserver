package application

import (
	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
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

	if !i.repository.CheckSessionActive(input.SessionToken) {
		return OutputData{Err: errors.SessionNotActive()}
	}

	i.repository.Disconnect(input.SessionToken)

	return OutputData{Status: i.repository.GetSessionStatus(input.SessionToken)}
}
