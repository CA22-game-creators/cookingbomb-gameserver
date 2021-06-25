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

func TestDiconnect(t *testing.T) {

	session.ClearSession()
	ctx := context.TODO()

	tests := []struct {
		title   string
		input   *pb.ConnectionRequest
		connect bool
		expect1 *pb.ConnectionResponse
		expect2 error
	}{
		{
			title: "[正常]正常なトークンで接続し、切断",
			input: &pb.ConnectionRequest{
				SessionToken: token.SessionToken.Valid,
			},
			connect: true,
			expect1: &pb.ConnectionResponse{
				Status: pb.ConnectionStatusEnum_DISCONNECTED_BY_CLIENT,
			},
			expect2: nil,
		},
		{
			title: "[異常]未接続のトークンで切断",
			input: &pb.ConnectionRequest{
				SessionToken: token.SessionToken.Valid,
			},
			connect: false,
			expect1: nil,
			expect2: errors.SessionNotActive(), // TODO: 本来バリデーションで弾かれる errors.InvalidArgument
		},
	}

	gs := controller.GameService{}

	for _, td := range tests {
		td := td

		t.Run(td.title, func(t *testing.T) {
			session.ClearSession()
			if td.connect {
				gs.Connect(ctx, td.input)
			}
			output1, output2 := gs.Disconnect(ctx, td.input)

			assert.Equal(t, td.expect1, output1)
			assert.Equal(t, td.expect2, output2)
		})
	}
}
