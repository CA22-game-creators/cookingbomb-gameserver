package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	connect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/connect"
	disconnect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/disconnect"
	getstatus "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/get_connection_status"
	mockDomain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/domain/model/account"
)

type testHandler struct {
	connect             connect.InputPort
	disconnect          disconnect.InputPort
	getConnectionStatus getstatus.InputPort

	repository *mockDomain.MockRepository
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)

	h.repository = mockDomain.NewMockRepository(ctrl)
	h.connect = connect.New(h.repository)
	h.disconnect = disconnect.New(h.repository)
	h.getConnectionStatus = getstatus.New(h.repository)
}
