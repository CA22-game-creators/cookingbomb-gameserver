package auth

import (
	api "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/api"
	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/session"
)

// トークンを認証する
func AuthToken(token string) (bool, error) {
	//TODO: API叩いてチェック
	account, err := api.GetAccountInfo(token)

	if err != nil {
		return false, err
	}

	id := account.Id
	// TODO: idがホワイトリストに居るか(ゲームに接続できるユーザーか？)
	if id == "" {
		return false, nil
	}

	return true, nil
}

// トークンを検証する
func CheckToken(token string) bool {
	return session.CheckSessionActive(token)
}
