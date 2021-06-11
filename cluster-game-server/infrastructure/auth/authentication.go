package auth

import (
	// api "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/api"
	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/session"
)

// トークンを認証する
func AuthToken(token string) (bool, error) {
	//TODO: API叩いてチェック

	//TODO: トークンをセッション処理関数に投げる
	session.ActivateSession(token)
	return true, nil
}

// トークンを検証する
func CheckToken(token string) bool {
	return session.CheckSessionActive(token)
}
