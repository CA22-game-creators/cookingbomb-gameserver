package account_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	account "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/model/account"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
)

func TestGetStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    pb.AccountInfo
		expected account.Account
	}{
		{
			input: pb.AccountInfo{
				Id:   "550e8400-e29b-41d4-a716-446655440000",
				Name: "Test User",
			},
			expected: account.Account{
				Id:   "550e8400-e29b-41d4-a716-446655440000",
				Name: "Test User",
			},
		},
		{
			input: pb.AccountInfo{
				Id:   "00000000-0000-0000-0000-000000000000",
				Name: "日本語",
			},
			expected: account.Account{
				Id:   "00000000-0000-0000-0000-000000000000",
				Name: "日本語",
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewFromAPIResponce: "+td.input.String(), func(t *testing.T) {
			t.Parallel()

			output := account.NewFromAPIResponce(&td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}
