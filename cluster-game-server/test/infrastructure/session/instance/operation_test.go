package session_instance_test

import (
	"testing"
	"time"

	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/session"
	session_model "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/model/session"
	token "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/token"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {

	session.ClearSession()

	instance := session.New()

	s := session_model.Session{
		Status: pb.ConnectionStatusEnum_CONNECTED,
	}

	s2 := session_model.Session{
		Status: pb.ConnectionStatusEnum_DISCONNECTED,
	}

	tests := []struct {
		title  string
		before func()
		test   func(*testing.T)
		after  func()
	}{
		{
			title:  "[正常]データを登録",
			before: func() {},
			test: func(t *testing.T) {
				instance.SetValue(token.SessionToken.Valid, s)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[正常]データを登録し、取得",
			before: func() {
				instance.SetValue(token.SessionToken.Valid, s)
			},
			test: func(t *testing.T) {
				output1, output2 := instance.GetValue(token.SessionToken.Valid)
				assert.Equal(t, s, output1)
				assert.True(t, output2)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[正常]データを登録し、期限を含めて取得",
			before: func() {
				instance.SetValue(token.SessionToken.Valid, s)
			},
			test: func(t *testing.T) {
				output1, output2, output3 := instance.GetValueWithEcpiration(token.SessionToken.Valid)
				assert.Equal(t, s, output1)
				assert.Equal(t, time.Time{}, output2)
				assert.True(t, output3)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[正常]データを期限付きで登録し、取得",
			before: func() {
				instance.SetValueWithExpiration(token.SessionToken.Valid, s)
			},
			test: func(t *testing.T) {
				output1, output2 := instance.GetValue(token.SessionToken.Valid)
				assert.Equal(t, s, output1)
				assert.True(t, output2)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[正常]データを期限付きで登録し、期限を含めて取得",
			before: func() {
				instance.SetValueWithExpiration(token.SessionToken.Valid, s)
			},
			test: func(t *testing.T) {
				output1, output2, output3 := instance.GetValueWithEcpiration(token.SessionToken.Valid)
				assert.Equal(t, output1, s)
				dur := time.Until(output2)
				assert.GreaterOrEqual(t, dur, time.Minute*25)
				assert.LessOrEqual(t, dur, time.Minute*30)
				assert.True(t, output3)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[正常]データを登録し、新たなデータを上書きして取得",
			before: func() {
				instance.SetValue(token.SessionToken.Valid, s)
				instance.SetValue(token.SessionToken.Valid, s2)
			},
			test: func(t *testing.T) {
				output1, output2, output3 := instance.GetValueWithEcpiration(token.SessionToken.Valid)
				assert.Equal(t, s2, output1)
				assert.Equal(t, time.Time{}, output2)
				assert.True(t, output3)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[正常]データを期限なしで登録し、新たなデータを期限付きで上書きして取得",
			before: func() {
				instance.SetValue(token.SessionToken.Valid, s)
				instance.SetValueWithExpiration(token.SessionToken.Valid, s2)
			},
			test: func(t *testing.T) {
				output1, output2, output3 := instance.GetValueWithEcpiration(token.SessionToken.Valid)
				assert.Equal(t, s2, output1)
				dur := time.Until(output2)
				assert.GreaterOrEqual(t, dur, time.Minute*25)
				assert.LessOrEqual(t, dur, time.Minute*30)
				assert.True(t, output3)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[正常]データを期限付きで登録し、新たなデータを期限なしで上書きして取得",
			before: func() {
				instance.SetValueWithExpiration(token.SessionToken.Valid, s)
				instance.SetValue(token.SessionToken.Valid, s2)
			},
			test: func(t *testing.T) {
				output1, output2, output3 := instance.GetValueWithEcpiration(token.SessionToken.Valid)
				assert.Equal(t, s2, output1)
				assert.Equal(t, time.Time{}, output2)
				assert.True(t, output3)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[異常]存在しないデータを取得",
			before: func() {
			},
			test: func(t *testing.T) {
				output1, output2 := instance.GetValue(token.SessionToken.Valid)
				assert.Equal(t, session_model.Session{}, output1)
				assert.False(t, output2)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[異常]存在しないデータを期限付きで取得",
			before: func() {
			},
			test: func(t *testing.T) {
				output1, output2, output3 := instance.GetValueWithEcpiration(token.SessionToken.Valid)
				assert.Equal(t, session_model.Session{}, output1)
				assert.Equal(t, time.Time{}, output2)
				assert.False(t, output3)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[異常]登録したデータをFlush後に取得",
			before: func() {
				instance.SetValue(token.SessionToken.Valid, s)
				instance.Flush()
			},
			test: func(t *testing.T) {
				output1, output2 := instance.GetValue(token.SessionToken.Valid)
				assert.Equal(t, session_model.Session{}, output1)
				assert.False(t, output2)
			},
			after: func() {
				instance.Flush()
			},
		},
		{
			title: "[異常]登録したデータをFlush後に期限付きで取得",
			before: func() {
				instance.SetValue(token.SessionToken.Valid, s)
				instance.Flush()
			},
			test: func(t *testing.T) {
				output1, output2, output3 := instance.GetValueWithEcpiration(token.SessionToken.Valid)
				assert.Equal(t, session_model.Session{}, output1)
				assert.Equal(t, time.Time{}, output2)
				assert.False(t, output3)
			},
			after: func() {
				instance.Flush()
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run(td.title, func(t *testing.T) {
			td.before()
			td.test(t)
			td.after()
		})
	}
}
