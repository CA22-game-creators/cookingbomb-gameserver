package session

import (
	"time"

	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/model/session"
	"github.com/patrickmn/go-cache"
)

type Instance struct {
	cache *cache.Cache
}

func New() *Instance {
	var instance = &Instance{
		cache: cache.New(30*time.Minute, 30*time.Second),
	}
	return instance
}

func (instance Instance) GetValue(token string) (session.Session, bool) {
	cv, found := instance.cache.Get(token)
	if found {
		return cv.(session.Session), true
	}
	return session.Session{}, false
}

func (instance Instance) GetValueWithEcpiration(token string) (session.Session, time.Time, bool) {
	cv, t, found := instance.cache.GetWithExpiration(token)
	if found {
		return cv.(session.Session), t, true
	}
	return session.Session{}, t, false
}

func (instance Instance) SetValue(token string, value session.Session) {
	instance.cache.Set(token, value, cache.NoExpiration)
}

func (instance Instance) SetValueWithExpiration(token string, value session.Session) {
	instance.cache.Set(token, value, cache.DefaultExpiration)
}

func (instance Instance) Flush() {
	instance.cache.Flush()
}
