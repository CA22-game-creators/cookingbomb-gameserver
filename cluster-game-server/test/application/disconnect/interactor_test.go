package interactor_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	disconnect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/disconnect"
	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	mockDomain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/domain/model/account"
	testdata "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/token"
)

type testHandler struct {
	disconnect disconnect.InputPort

	repository *mockDomain.MockRepository
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)

	h.repository = mockDomain.NewMockRepository(ctrl)
	h.disconnect = disconnect.New(h.repository)
}
func TestDisconnect(t *testing.T) {
	t.Parallel()
	var status domain.StatusEnum

	tests := []struct {
		title    string
		before   func(testHandler)
		input    disconnect.InputData
		expected disconnect.OutputData
	}{
		{
			title: "[正常]接続済みを想定したセッショントークンを処理",
			before: func(h testHandler) {
				status = domain.CONNECTED
				h.repository.EXPECT().Disconnect(testdata.SessionToken.Valid).Do(func(_ string) interface{} {
					status = domain.DISCONNECTED_BY_CLIENT
					return nil
				})
				h.repository.EXPECT().GetSessionStatus(testdata.SessionToken.Valid).DoAndReturn(func(_ string) interface{} {
					return status
				}).AnyTimes()
			},
			input:    disconnect.InputData{SessionToken: testdata.SessionToken.Valid},
			expected: disconnect.OutputData{Status: domain.DISCONNECTED_BY_CLIENT},
		},
		{
			title: "[異常]未接続を想定したセッショントークンを処理",
			before: func(h testHandler) {
				status = domain.UNSPECIFIED
				h.repository.EXPECT().GetSessionStatus(testdata.SessionToken.Valid).DoAndReturn(func(_ string) interface{} {
					return status
				}).AnyTimes()
			},
			input:    disconnect.InputData{SessionToken: testdata.SessionToken.Valid},
			expected: disconnect.OutputData{Err: errors.SessionNotActive()},
		},
		{
			title: "[異常]切断済みを想定したセッショントークンを処理",
			before: func(h testHandler) {
				status = domain.DISCONNECTED_BY_CLIENT
				h.repository.EXPECT().GetSessionStatus(testdata.SessionToken.Valid).DoAndReturn(func(_ string) interface{} {
					return status
				}).AnyTimes()
			},
			input:    disconnect.InputData{SessionToken: testdata.SessionToken.Valid},
			expected: disconnect.OutputData{Err: errors.SessionNotActive()},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("application/disconnect:"+td.title, func(t *testing.T) {
			var tester testHandler
			tester.setupTest(t)
			if td.before != nil {
				td.before(tester)
			}

			actual := tester.disconnect.Handle(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}
