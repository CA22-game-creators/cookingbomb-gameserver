package application_get_connection_status_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	getstatus "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/get_connection_status"
	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	mockDomain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/domain/model/account"
	testdata "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/token"
)

type testHandler struct {
	get_connection_status getstatus.InputPort

	repository *mockDomain.MockRepository
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)

	h.repository = mockDomain.NewMockRepository(ctrl)
	h.get_connection_status = getstatus.New(h.repository)
}

func TestDisconnect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		before   func(testHandler)
		input    getstatus.InputData
		expected getstatus.OutputData
	}{
		{
			title: "[正常]接続済みを想定したセッショントークンを処理",
			before: func(h testHandler) {
				h.repository.EXPECT().GetSessionStatus(testdata.SessionToken.Valid).Return(domain.CONNECTED)
			},
			input:    getstatus.InputData{SessionToken: testdata.SessionToken.Valid},
			expected: getstatus.OutputData{Status: domain.CONNECTED},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("application/get_session_status:"+td.title, func(t *testing.T) {
			var tester testHandler
			tester.setupTest(t)
			if td.before != nil {
				td.before(tester)
			}

			actual := tester.get_connection_status.Handle(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}
