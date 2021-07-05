package application

import (
	"errors"

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

	_, err := i.repository.Find(input.SessionToken)
	if err != nil {
		return OutputData{Err: err}
	}

	if i.repository.CheckSessionActive(input.SessionToken) {
		return OutputData{Err: errors.New("already connected")}
	}

	i.repository.Connect(input.SessionToken)

	return OutputData{Status: i.repository.GetSessionStatus(input.SessionToken)}
}
