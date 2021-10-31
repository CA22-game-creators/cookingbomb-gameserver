package presentation_test

import (
	"context"
	"testing"

	mockGameDataStream "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/application/game_data_stream"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/golang/mock/gomock"

	controller "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/presentation"
)

type testHandler struct {
	controller pb.GameServicesServer

	context        context.Context
	gameDataStream *mockGameDataStream.MockInputPort
}

func (h *testHandler) setupTest(t *testing.T) {
	h.context = context.TODO()

	ctrl := gomock.NewController(t)
	h.gameDataStream = mockGameDataStream.NewMockInputPort(ctrl)

	h.controller = controller.New(h.gameDataStream)
}
