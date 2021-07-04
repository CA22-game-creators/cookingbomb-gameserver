// +build wireinject

package di

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/google/wire"

	connect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/connect"
	disconnect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/disconnect"
	getConnectionStatus "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/get_connection_status"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/cache"
	repository "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/repository"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/presentation"
)

func DI() pb.GameServicesServer {
	wire.Build(
		presentation.New,
		disconnect.New,
		connect.New,
		getConnectionStatus.New,
		repository.New,
		cache.New,
	)

	return nil
}
