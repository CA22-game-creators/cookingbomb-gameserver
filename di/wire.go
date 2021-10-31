//go:build wireinject
// +build wireinject

package di

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/google/wire"

	gameDataStream "github.com/CA22-game-creators/cookingbomb-gameserver/application/game_data_stream"
	account "github.com/CA22-game-creators/cookingbomb-gameserver/infrastructure/repository/account"
	character "github.com/CA22-game-creators/cookingbomb-gameserver/infrastructure/repository/character"
	"github.com/CA22-game-creators/cookingbomb-gameserver/presentation"
)

func DI() pb.GameServicesServer {
	wire.Build(
		presentation.New,
		gameDataStream.New,
		account.New,
		character.New,
	)

	return nil
}
