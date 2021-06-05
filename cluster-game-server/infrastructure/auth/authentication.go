package auth

import (
	c "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/cache"
)

var cacheInstance *c.TokenCacheInstance = c.New()

type User struct {
	id   string
	name string
}

func CheckAuthToken(token string) bool {
	cache := checkAuthTokenCached(token)
	if cache {
		return true
	}

	//TODO
	return true
}

func checkAuthTokenCached(token string) bool {
	found := cacheInstance.CheckTokenCached(token)
	return found
}
