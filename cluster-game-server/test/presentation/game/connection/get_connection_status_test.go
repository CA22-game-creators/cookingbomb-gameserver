package connection_controller_test

import (
	"context"
	"testing"

	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/session"
	controller "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/presentation/game"
	token "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/token"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/stretchr/testify/assert"
)

func TestGetConnectionStatus(t *testing.T) {

	session.ClearSession()
	ctx := context.TODO()

	gs := controller.GameService{}

	tests := []struct {
		title   string
		input   *pb.ConnectionRequest
		before  func(*testing.T, *pb.ConnectionRequest)
		expect1 *pb.ConnectionResponse
		expect2 error
	}{
		{
			title: "[正常]正常なトークンで接続し、ステータスを取得",
			input: &pb.ConnectionRequest{
				SessionToken: token.SessionToken.Valid,
			},
			before: func(t *testing.T, input *pb.ConnectionRequest) {
				_, _ = gs.Connect(ctx, input)
			},
			expect1: &pb.ConnectionResponse{
				Status: pb.ConnectionStatusEnum_CONNECTED,
			},
			expect2: nil,
		},
		{
			title: "[異常]未接続の正常なトークンでステータスを取得",
			input: &pb.ConnectionRequest{
				SessionToken: token.SessionToken.Valid,
			},
			before: func(t *testing.T, input *pb.ConnectionRequest) {
			},
			expect1: nil,
			expect2: errors.NoStatusFound(),
		},
		{
			title: "[異常]不正なトークンでステータスを取得",
			input: &pb.ConnectionRequest{
				SessionToken: token.SessionToken.Invalid,
			},
			before: func(t *testing.T, input *pb.ConnectionRequest) {
			},
			expect1: nil,
			expect2: errors.NoStatusFound(), // 本来はinvalidargument
		},
	}

	for _, td := range tests {
		td := td

		t.Run(td.title, func(t *testing.T) {
			session.ClearSession()

			td.before(t, td.input)
			output1, output2 := gs.GetConnectionStatus(ctx, td.input)

			assert.Equal(t, td.expect1, output1)
			assert.Equal(t, td.expect2, output2)
		})
	}
}
