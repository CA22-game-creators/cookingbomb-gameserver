package application_game_data_stream_test

import (
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gameDataStream "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/game_data_stream"
	account "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	character "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/character"
	charactertd "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/character"
	stream "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/stream"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

func TestGameDataStreamHandle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		before   func(testHandler, *stream.DummyStream)
		expected gameDataStream.OutputData
		check    func(*testing.T, *stream.DummyStream)
	}{
		{
			title: "[正常]正常にストリームが終了する場合",
			before: func(h testHandler, s *stream.DummyStream) {
				s.RecvExpect = nil
				h.characterrepo.EXPECT().GetAll().Return(nil).AnyTimes()
			},
			expected: gameDataStream.OutputData{
				Err: nil,
			},
		},
		{
			title: "[正常]データを受信し、正常に書き込まれること",
			before: func(h testHandler, s *stream.DummyStream) {
				s.RecvExpect = append(s.RecvExpect, &pb.GameDataRequest{
					SessionToken: "Test",
					Message: &pb.GameDataRequest_CharacterUpdate{
						CharacterUpdate: &charactertd.Character,
					},
				})
				h.characterrepo.EXPECT().GetAll().Return(nil).AnyTimes()
				h.characterrepo.EXPECT().Update(&charactertd.Character).Times(1)
				h.accountrepo.EXPECT().GetSessionStatus("Test").Return(account.CONNECTED).AnyTimes()
			},
			expected: gameDataStream.OutputData{
				Err: nil,
			},
		},
		{
			title: "[正常]データを取得し、正常に送信されること",
			before: func(h testHandler, s *stream.DummyStream) {
				s.RecvExpect = append(s.RecvExpect, &pb.GameDataRequest{
					SessionToken: "Test",
					Message: &pb.GameDataRequest_CharacterUpdate{
						CharacterUpdate: &charactertd.Character,
					},
				})
				s.RecvDelay = (time.Second)
				h.characterrepo.EXPECT().GetAll().Return(&[]character.Character{
					&charactertd.Character,
				}).AnyTimes()
				h.characterrepo.EXPECT().Update(&charactertd.Character).AnyTimes()
				h.accountrepo.EXPECT().GetSessionStatus("Test").Return(account.CONNECTED).AnyTimes()
			},
			expected: gameDataStream.OutputData{
				Err: nil,
			},
			check: func(t *testing.T, s *stream.DummyStream) {
				assert.NotEmpty(t, s.SendLog)
				assert.Equal(t, &charactertd.Character, s.SendLog[0].Message.(*pb.GameDataResponse_CharacterDatas).CharacterDatas.Characters[0])
			},
		},
		{
			title: "[異常]ストリームが異常終了する場合",
			before: func(h testHandler, s *stream.DummyStream) {
				s.RecvExpect = nil
				s.RecvErr = status.Errorf(codes.Internal, io.ErrUnexpectedEOF.Error())
				h.characterrepo.EXPECT().GetAll().Return(nil).AnyTimes()
			},
			expected: gameDataStream.OutputData{
				Err: status.Errorf(codes.Internal, io.ErrUnexpectedEOF.Error()),
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("application/game_data_stream-handle:"+td.title, func(t *testing.T) {
			var tester testHandler
			tester.setupTest(t)
			s := stream.DummyStream{}
			if td.before != nil {
				td.before(tester, &s)
			}

			input := gameDataStream.InputData{
				Stream: &s,
			}
			actual := tester.gameDataStream.Handle(input)
			assert.Equal(t, td.expected, actual)

			if td.check != nil {
				td.check(t, &s)
			}
		})
	}
}
