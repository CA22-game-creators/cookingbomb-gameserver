package auth

import (
	c "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/cache"
)

var cacheInstance *c.TokenCacheInstance = c.New()

// トークンを認証する
func AuthToken(token string) (bool, error) {
	//TODO: API叩いてチェック

	//TODO: トークンをセッション処理関数に投げる

	return true, nil
}

// トークンを検証する
func CheckToken(token string) (bool, error) {
	cache := checkTokenCached(token)
	if cache {
		return true, nil
	}
	return AuthToken(token)
}

func checkTokenCached(token string) bool {
	found := cacheInstance.CheckTokenCached(token)
	return found
}
