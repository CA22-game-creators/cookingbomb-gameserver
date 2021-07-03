package presentation_test

import (
	"testing"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	connect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/connect"
	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	testdata "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/token"
)

func TestConnect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		before    func(testHandler)
		input     *pb.ConnectionRequest
		expected1 *pb.ConnectionResponse
		expected2 error
	}{
		{
			title: "【正常系】セッショントークンからステータスをCONNECTEDにできる",
			before: func(h testHandler) {
				input := connect.InputData{SessionToken: testdata.SessionToken.Valid}
				output := connect.OutputData{Status: domain.CONNECTED}
				h.connect.EXPECT().Handle(input).Return(output)
			},
			input:     &pb.ConnectionRequest{SessionToken: testdata.SessionToken.Valid},
			expected1: &pb.ConnectionResponse{Status: pb.ConnectionStatusEnum_CONNECTED},
			expected2: nil,
		},
		{
			title:     "【異常系】セッショントークンが不整値",
			input:     &pb.ConnectionRequest{SessionToken: testdata.SessionToken.Invalid},
			expected1: nil,
			expected2: status.Error(codes.InvalidArgument, "sessionTokenが不正な形式です"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Connect:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			if td.before != nil {
				td.before(tester)
			}

			actual1, actual2 := tester.controller.Connect(tester.context, td.input)
			assert.Equal(t, td.expected1, actual1)
			assert.Equal(t, td.expected2, actual2)
		})
	}
}
