package application_connect_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	connect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/connect"
	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	mockDomain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/mock/domain/model/account"
	testdata "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/token"
)

type testHandler struct {
	connect connect.InputPort

	repository *mockDomain.MockRepository
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)

	h.repository = mockDomain.NewMockRepository(ctrl)
	h.connect = connect.New(h.repository)
}

func TestConnect(t *testing.T) {
	t.Parallel()
	var status domain.StatusEnum

	tests := []struct {
		title    string
		before   func(testHandler)
		input    connect.InputData
		expected connect.OutputData
	}{
		{
			title: "[正常]正常なセッショントークンを処理",
			before: func(h testHandler) {
				status = domain.UNSPECIFIED
				h.repository.EXPECT().Find(testdata.SessionToken.Valid).Return(testdata.Account, nil)
				h.repository.EXPECT().Connect(testdata.SessionToken.Valid).Do(func(_ string) interface{} {
					status = domain.CONNECTED
					return nil
				})
				h.repository.EXPECT().GetSessionStatus(testdata.SessionToken.Valid).DoAndReturn(func(_ string) interface{} {
					return status
				}).AnyTimes()
			},
			input:    connect.InputData{SessionToken: testdata.SessionToken.Valid},
			expected: connect.OutputData{Status: domain.CONNECTED},
		},
		{
			title: "[正常]切断済みを想定したセッショントークンを処理",
			before: func(h testHandler) {
				status = domain.DISCONNECTED_BY_CLIENT
				h.repository.EXPECT().Find(testdata.SessionToken.Valid).Return(testdata.Account, nil)
				h.repository.EXPECT().Connect(testdata.SessionToken.Valid).Do(func(_ string) interface{} {
					status = domain.CONNECTED
					return nil
				})
				h.repository.EXPECT().GetSessionStatus(testdata.SessionToken.Valid).DoAndReturn(func(_ string) interface{} {
					return status
				}).AnyTimes()
			},
			input:    connect.InputData{SessionToken: testdata.SessionToken.Valid},
			expected: connect.OutputData{Status: domain.CONNECTED},
		},
		{
			title: "[異常]接続済みを想定したセッショントークンを処理",
			before: func(h testHandler) {
				status = domain.CONNECTED
				h.repository.EXPECT().Find(testdata.SessionToken.Valid).Return(testdata.Account, nil)
				h.repository.EXPECT().GetSessionStatus(testdata.SessionToken.Valid).DoAndReturn(func(_ string) interface{} {
					return status
				}).AnyTimes()
			},
			input:    connect.InputData{SessionToken: testdata.SessionToken.Valid},
			expected: connect.OutputData{Err: errors.InvalidOperation()},
		},
		{
			title: "[異常]無効なセッショントークンを処理",
			before: func(h testHandler) {
				status = domain.UNSPECIFIED
				h.repository.EXPECT().Find(testdata.SessionToken.Invalid).Return(domain.Account{}, errors.AuthAPIThrowError("test"))
			},
			input:    connect.InputData{SessionToken: testdata.SessionToken.Invalid},
			expected: connect.OutputData{Err: errors.AuthAPIThrowError("test")},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("application/connect:"+td.title, func(t *testing.T) {
			var tester testHandler
			tester.setupTest(t)
			if td.before != nil {
				td.before(tester)
			}

			actual := tester.connect.Handle(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}
