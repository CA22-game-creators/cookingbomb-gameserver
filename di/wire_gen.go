// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/CA22-game-creators/cookingbomb-gameserver/application/game_data_stream"
	"github.com/CA22-game-creators/cookingbomb-gameserver/infrastructure/repository/account"
	infra2 "github.com/CA22-game-creators/cookingbomb-gameserver/infrastructure/repository/character"
	"github.com/CA22-game-creators/cookingbomb-gameserver/presentation"
	"github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

// Injectors from wire.go:

func DI() game.GameServicesServer {
	repository := infra.New()
	domainRepository := infra2.New()
	inputPort := application.New(repository, domainRepository)
	gameServicesServer := presentation.New(inputPort)
	return gameServicesServer
}
