package auth

import (
	api "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/api"
)

// トークンを認証する
func CheckTokenPermission(token string) (bool, error) {
	//TODO: API叩いてチェック
	account, err := api.GetAccountInfo(token)

	if err != nil {
		return false, err
	}

	id := account.ID
	// TODO: idがホワイトリストに居るか(ゲームに接続できるユーザーか？)
	if id == "" {
		return false, nil
	}

	return true, nil
}
