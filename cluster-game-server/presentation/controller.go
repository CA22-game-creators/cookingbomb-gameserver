package presentation

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"

	connect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/connect"
	disconnect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/disconnect"
	getConnectionStatus "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/get_connection_status"
	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
)

type controller struct {
	connect             connect.InputPort
	disconnect          disconnect.InputPort
	getConnectionStatus getConnectionStatus.InputPort
}

func New(co connect.InputPort, di disconnect.InputPort, gc getConnectionStatus.InputPort) pb.GameServicesServer {
	return &controller{
		connect:             co,
		disconnect:          di,
		getConnectionStatus: gc,
	}
}

func StatusMapper(status domain.StatusEnum) pb.ConnectionStatusEnum {
	switch status {
	case domain.CONNECTED:
		return pb.ConnectionStatusEnum_CONNECTED
	case domain.DISCONNECTED_BY_CLIENT:
		return pb.ConnectionStatusEnum_DISCONNECTED_BY_CLIENT
	default:
		return pb.ConnectionStatusEnum_CONNECTION_UNSPECIFIED
	}
}
