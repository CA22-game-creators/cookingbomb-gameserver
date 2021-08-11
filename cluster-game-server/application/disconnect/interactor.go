package application

import (
	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return OutputData{
		Err: status.Error(codes.Unimplemented, "duplicated"),
	}
	/*
		if status := i.repository.GetSessionStatus(input.SessionToken); !status.IsActive() {
			return OutputData{Err: errors.SessionNotActive()}
		}

		i.repository.Disconnect(input.SessionToken)

		return OutputData{Status: i.repository.GetSessionStatus(input.SessionToken)}
	*/
}
