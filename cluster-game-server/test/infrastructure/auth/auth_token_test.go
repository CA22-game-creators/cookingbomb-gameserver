package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	errors "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	auth "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/auth"
	token "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/token"
)

func TestAuthToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input     string
		expected1 bool
		expected2 error
	}{
		{
			title:     "[正常]正常なトークンの認証",
			input:     token.SessionToken.Valid,
			expected1: true,
			expected2: nil,
		},
		{
			title:     "[異常]無効なトークンの認証",
			input:     token.SessionToken.Invalid,
			expected1: false,
			expected2: errors.AuthAPIThrowError(),
		},
	}

	for _, td := range tests {
		td := td

		t.Run(td.title, func(t *testing.T) {
			t.Parallel()

			output1, output2 := auth.AuthToken(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
