package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	auth "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/auth"
	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/session"
	token "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/token"
)

func TestCheckToken(t *testing.T) {

	//並列不可
	session.ClearSession()

	tests := []struct {
		title    string
		before   func()
		after    func()
		input    string
		expected bool
	}{
		{
			title: "[正常]正常なトークンの確認",
			before: func() {
				auth.AuthToken(token.SessionToken.Valid)
			},
			after: func() {
				session.ClearSession()
			},
			input:    token.SessionToken.Valid,
			expected: true,
		},
		{
			title: "[異常]無効なトークンの確認",
			before: func() {
			},
			after: func() {
				session.ClearSession()
			},
			input:    token.SessionToken.Invalid,
			expected: false,
		},
		{
			title: "[異常]未認証のトークンの確認",
			before: func() {
			},
			after: func() {
				session.ClearSession()
			},
			input:    token.SessionToken.Valid2,
			expected: false,
		},
		{
			title: "[異常]切断済みのトークンの確認",
			before: func() {
				auth.AuthToken(token.SessionToken.Valid3)
				session.ForceEndSession(token.SessionToken.Valid3)
			},
			after: func() {
				session.ClearSession()
			},
			input:    token.SessionToken.Valid3,
			expected: false,
		},
	}

	for _, td := range tests {
		td := td

		t.Run(td.title, func(t *testing.T) {
			td.before()
			output := auth.CheckToken(td.input)

			assert.Equal(t, td.expected, output)
			td.after()
		})
	}
}
