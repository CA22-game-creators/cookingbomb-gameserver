package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Instance *cache.Cache = cache.New(30*time.Minute, 30*time.Second)
