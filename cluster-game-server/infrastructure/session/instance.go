package session

import (
	"time"

	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/model/session"
	"github.com/patrickmn/go-cache"
)

type SessionInstance struct {
	cache *cache.Cache
}

func New() *SessionInstance {
	instance := &SessionInstance{
		cache: cache.New(30*time.Minute, 30*time.Second),
	}
	return instance
}

func (instance SessionInstance) GetValue(token string) (session.Session, bool) {
	cv, found := instance.cache.Get(token)
	if found {
		return cv.(session.Session), true
	}
	return session.Session{}, false
}

func (instance SessionInstance) SetValue(token string, value session.Session) {
	instance.cache.Set(token, value, cache.NoExpiration)
}

func (instance SessionInstance) SetValueWithExpiration(token string, value session.Session) {
	instance.cache.Set(token, value, cache.DefaultExpiration)
}
