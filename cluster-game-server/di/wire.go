// +build wireinject

package di

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/google/wire"

	connect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/connect"
	disconnect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/disconnect"
	gameDataStream "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/game_data_stream"
	getConnectionStatus "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/get_connection_status"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/cache"
	account "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/repository/account"
	character "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/repository/character"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/presentation"
)

func DI() pb.GameServicesServer {
	wire.Build(
		presentation.New,
		disconnect.New,
		connect.New,
		getConnectionStatus.New,
		gameDataStream.New,
		account.New,
		character.New,
		cache.New,
	)

	return nil
}
