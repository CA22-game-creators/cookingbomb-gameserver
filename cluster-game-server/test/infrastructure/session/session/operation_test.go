package session_test

import (
	"testing"

	errors "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/session"
	token "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/token"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {

	session.ClearSession()

	tests := []struct {
		title        string
		before       func()
		test         func(t *testing.T)
		after        func()
		token        string
		expectStatus pb.ConnectionStatusEnum
		expectActive bool
	}{
		{
			title: "[正常]セッションを有効化",
			before: func() {
				_ = session.ActivateSession(token.SessionToken.Valid)
			},
			after: func() {
				session.ClearSession()
			},
			token:        token.SessionToken.Valid,
			expectStatus: pb.ConnectionStatusEnum_CONNECTED,
			expectActive: true,
		},
		{
			title: "[正常]セッションを有効化後サーバーにより切断",
			before: func() {
				_ = session.ActivateSession(token.SessionToken.Valid)
			},
			test: func(t *testing.T) {
				err := session.EndSessionByServer(token.SessionToken.Valid)
				assert.Nil(t, err)
			},
			after: func() {
				session.ClearSession()
			},
			token:        token.SessionToken.Valid,
			expectStatus: pb.ConnectionStatusEnum_DISCONNECTED_BY_SERVER,
			expectActive: false,
		},
		{
			title: "[正常]セッションを有効化後クライアントにより切断",
			before: func() {
				_ = session.ActivateSession(token.SessionToken.Valid)
			},
			test: func(t *testing.T) {
				err := session.EndSessionByClient(token.SessionToken.Valid)
				assert.Nil(t, err)
			},
			after: func() {
				session.ClearSession()
			},
			token:        token.SessionToken.Valid,
			expectStatus: pb.ConnectionStatusEnum_DISCONNECTED_BY_CLIENT,
			expectActive: false,
		},
		{
			title: "[正常]セッションを有効化後強制的に切断",
			before: func() {
				_ = session.ActivateSession(token.SessionToken.Valid)
				session.ForceEndSession(token.SessionToken.Valid)
			},
			after: func() {
				session.ClearSession()
			},
			token:        token.SessionToken.Valid,
			expectStatus: pb.ConnectionStatusEnum_DISCONNECTED,
			expectActive: false,
		},
		{
			title: "[正常]セッションを有効化後強制切断し、再度接続",
			before: func() {
				_ = session.ActivateSession(token.SessionToken.Valid)
				session.ForceEndSession(token.SessionToken.Valid)
				_ = session.ActivateSession(token.SessionToken.Valid)
			},
			after: func() {
				session.ClearSession()
			},
			token:        token.SessionToken.Valid,
			expectStatus: pb.ConnectionStatusEnum_CONNECTED,
			expectActive: true,
		},
		{
			title: "[正常]セッションを有効化後、全セッション削除",
			before: func() {
				_ = session.ActivateSession(token.SessionToken.Valid)
				session.ClearSession()
			},
			after: func() {
				session.ClearSession()
			},
			token:        token.SessionToken.Valid,
			expectStatus: pb.ConnectionStatusEnum_CONNECTION_UNSPECIFIED,
			expectActive: false,
		},
		{
			title: "[正常]未処理のトークンを検証",
			before: func() {
			},
			after: func() {
				session.ClearSession()
			},
			token:        token.SessionToken.Valid,
			expectStatus: pb.ConnectionStatusEnum_CONNECTION_UNSPECIFIED,
			expectActive: false,
		},
		{
			title: "[異常]すでに接続済みのセッションを、再度有効化処理",
			before: func() {
				_ = session.ActivateSession(token.SessionToken.Valid)
			},
			test: func(t *testing.T) {
				err := session.ActivateSession(token.SessionToken.Valid)
				assert.ErrorIs(t, errors.InvalidOperation(), err)
			},
			after: func() {
				session.ClearSession()
			},
			token:        token.SessionToken.Valid,
			expectStatus: pb.ConnectionStatusEnum_CONNECTED,
			expectActive: true,
		},
		{
			title: "[異常]すでに切断済みのセッションを、再度切断処理",
			before: func() {
				_ = session.ActivateSession(token.SessionToken.Valid)
				session.ForceEndSession(token.SessionToken.Valid)
			},
			test: func(t *testing.T) {
				err1 := session.EndSessionByClient(token.SessionToken.Valid)
				err2 := session.EndSessionByServer(token.SessionToken.Valid)
				assert.ErrorIs(t, errors.InvalidOperation(), err1)
				assert.ErrorIs(t, errors.InvalidOperation(), err2)
			},
			after: func() {
				session.ClearSession()
			},
			token:        token.SessionToken.Valid,
			expectStatus: pb.ConnectionStatusEnum_DISCONNECTED,
			expectActive: false,
		},
		{
			title: "[異常]未処理のトークンに対して切断処理",
			before: func() {
			},
			test: func(t *testing.T) {
				err1 := session.EndSessionByClient(token.SessionToken.Valid)
				err2 := session.EndSessionByServer(token.SessionToken.Valid)
				assert.ErrorIs(t, errors.InvalidOperation(), err1)
				assert.ErrorIs(t, errors.InvalidOperation(), err2)
			},
			after: func() {
				session.ClearSession()
			},
			token:        token.SessionToken.Valid,
			expectStatus: pb.ConnectionStatusEnum_CONNECTION_UNSPECIFIED,
			expectActive: false,
		},
	}

	for _, td := range tests {
		td := td

		t.Run(td.title, func(t *testing.T) {
			td.before()

			if td.test != nil {
				td.test(t)
			}

			outputStatus := session.GetSessionStatus(td.token)
			outputActive := session.CheckSessionActive(td.token)

			assert.Equal(t, td.expectStatus, outputStatus)
			assert.Equal(t, td.expectActive, outputActive)

			td.after()
		})
	}
}
