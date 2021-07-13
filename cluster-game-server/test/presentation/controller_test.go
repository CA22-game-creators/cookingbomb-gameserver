package presentation_test

import (
	"context"
	"testing"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/golang/mock/gomock"

	mockConnect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/application/connect"
	mockDisconnect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/application/disconnect"
	mockGameDataStream "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/application/game_data_stream"
	mockGetConnectionStatus "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/application/get_connection_status"

	controller "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/presentation"
)

type testHandler struct {
	controller pb.GameServicesServer

	context             context.Context
	connect             *mockConnect.MockInputPort
	disconnect          *mockDisconnect.MockInputPort
	getConnectionStatus *mockGetConnectionStatus.MockInputPort
	gameDataStream      *mockGameDataStream.MockInputPort
}

func (h *testHandler) setupTest(t *testing.T) {
	h.context = context.TODO()

	ctrl := gomock.NewController(t)
	h.connect = mockConnect.NewMockInputPort(ctrl)
	h.disconnect = mockDisconnect.NewMockInputPort(ctrl)
	h.getConnectionStatus = mockGetConnectionStatus.NewMockInputPort(ctrl)
	h.gameDataStream = mockGameDataStream.NewMockInputPort(ctrl)

	h.controller = controller.New(h.connect, h.disconnect, h.getConnectionStatus, h.gameDataStream)
}
