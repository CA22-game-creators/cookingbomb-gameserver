// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/connect"
	application2 "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/disconnect"
	application4 "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/game_data_stream"
	application3 "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/get_connection_status"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/cache"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/repository"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/presentation"
	"github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

// Injectors from wire.go:

func DI() game.GameServicesServer {
	cacheCache := cache.New()
	repository := infra.New(cacheCache)
	inputPort := application.New(repository)
	applicationInputPort := application2.New(repository)
	inputPort2 := application3.New(repository)
	inputPort3 := application4.New(repository)
	gameServicesServer := presentation.New(inputPort, applicationInputPort, inputPort2, inputPort3)
	return gameServicesServer
}
