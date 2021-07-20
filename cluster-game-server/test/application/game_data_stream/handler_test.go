package application_game_data_stream_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	gameDataStream "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/game_data_stream"
	mockDomainAccount "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/domain/model/account"
	mockDomainCharacter "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/domain/model/character"
)

type testHandler struct {
	gameDataStream gameDataStream.InputPort

	accountrepo   *mockDomainAccount.MockRepository
	characterrepo *mockDomainCharacter.MockRepository
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)

	h.accountrepo = mockDomainAccount.NewMockRepository(ctrl)
	h.characterrepo = mockDomainCharacter.NewMockRepository(ctrl)
	h.gameDataStream = gameDataStream.New(h.accountrepo, h.characterrepo)
}
