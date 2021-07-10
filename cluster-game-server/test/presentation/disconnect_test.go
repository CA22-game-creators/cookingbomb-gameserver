package presentation_test

import (
	"testing"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/stretchr/testify/assert"

	disconnect "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/disconnect"
	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	testdata "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/token"
)

func TestDisonnect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		before    func(testHandler)
		input     *pb.ConnectionRequest
		expected1 *pb.ConnectionResponse
		expected2 error
	}{
		{
			title: "【正常系】セッショントークンからステータスをDISCONNECTEDにできる",
			before: func(h testHandler) {
				input := disconnect.InputData{SessionToken: testdata.SessionToken.Valid}
				output := disconnect.OutputData{Status: domain.DISCONNECTED_BY_CLIENT}
				h.disconnect.EXPECT().Handle(input).Return(output)
			},
			input:     &pb.ConnectionRequest{SessionToken: testdata.SessionToken.Valid},
			expected1: &pb.ConnectionResponse{Status: pb.ConnectionStatusEnum_DISCONNECTED_BY_CLIENT},
			expected2: nil,
		},
		{
			title:     "【異常系】セッショントークンが不整値",
			input:     &pb.ConnectionRequest{SessionToken: testdata.SessionToken.Invalid},
			expected1: nil,
			expected2: errors.InvalidArgument("sessionTokenが不正な形式です"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("presentation/Disonnect:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			if td.before != nil {
				td.before(tester)
			}

			actual1, actual2 := tester.controller.Disconnect(tester.context, td.input)
			assert.Equal(t, td.expected1, actual1)
			assert.Equal(t, td.expected2, actual2)
		})
	}
}
