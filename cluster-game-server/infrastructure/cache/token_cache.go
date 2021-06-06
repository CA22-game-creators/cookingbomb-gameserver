package cache

import (
	"time"

	user "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/model/user"
	"github.com/patrickmn/go-cache"
)

type TokenCacheInstance struct {
	cache *cache.Cache
}

func New() *TokenCacheInstance {
	instance := &TokenCacheInstance{
		cache: cache.New(30*time.Minute, 30*time.Second),
	}
	return instance
}

func (instance TokenCacheInstance) CheckTokenCached(token string) bool {
	_, found := instance.cache.Get(token)
	return found
}

func (instance TokenCacheInstance) GetValue(token string) (user.User, bool) {
	cv, found := instance.cache.Get(token)
	if found {
		return cv.(user.User), true
	}
	return user.User{}, false
}

func (instance TokenCacheInstance) SetValue(token string, value user.User) {
	instance.cache.Set(token, value, cache.NoExpiration)
}
