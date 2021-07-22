package presentation_test

import (
	"testing"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/stretchr/testify/assert"

	application "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/game_data_stream"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"

	testdata "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/stream"
)

func TestGameDataStream(t *testing.T) {
	t.Parallel()

	ds := &testdata.DummyStream{}

	tests := []struct {
		title    string
		before   func(testHandler)
		input    pb.GameServices_GameDataStreamServer
		expected error
	}{
		{
			title: "[正常] Streamがエラーが起きず終了した場合",
			before: func(h testHandler) {
				h.gameDataStream.EXPECT().Handle(application.InputData{
					Stream: ds,
				}).Return(application.OutputData{})
			},
			input:    ds,
			expected: nil,
		},
		{
			title: "[異常] Streamでエラーが発生し終了した場合",
			before: func(h testHandler) {
				h.gameDataStream.EXPECT().Handle(application.InputData{
					Stream: ds,
				}).Return(application.OutputData{
					Err: errors.SessionNotActive(),
				})
			},
			input:    ds,
			expected: errors.SessionNotActive(),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("presentation/GameDataStream:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			if td.before != nil {
				td.before(tester)
			}

			actual := tester.controller.GameDataStream(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}
