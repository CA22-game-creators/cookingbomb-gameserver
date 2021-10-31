package presentation_test

import (
	"testing"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/stretchr/testify/assert"

	application "github.com/CA22-game-creators/cookingbomb-gameserver/application/game_data_stream"
	td "github.com/CA22-game-creators/cookingbomb-gameserver/test/testdata/stream"
)

func TestGameDataStream(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		before   func(testHandler)
		input    pb.GameServices_GameDataStreamServer
		expected error
	}{
		{
			title: "【正常系】",
			before: func(h testHandler) {
				input := application.InputData{Stream: &td.Stream{}}
				h.gameDataStream.EXPECT().Handle(input).Return(application.OutputData{})
			},
			input:    &td.Stream{},
			expected: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("GameDataStream:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			td.before(tester)

			result := tester.controller.GameDataStream(td.input)
			assert.Equal(t, td.expected, result)
		})
	}
}
