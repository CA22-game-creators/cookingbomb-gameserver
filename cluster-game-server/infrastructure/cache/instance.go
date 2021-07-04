package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func New() *cache.Cache {
	return cache.New(30*time.Minute, 30*time.Second)
}
