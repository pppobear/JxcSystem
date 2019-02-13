package filters

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/config"
)

func RegisterSession() gin.HandlerFunc {
	store, _ := sessions.NewRedisStore(
		10,
		"tcp",
		config.GetEnv().RedisIP+":"+config.GetEnv().RedisPort,
		config.GetEnv().RedisPassword,
		[]byte("secret"),
	)
	return sessions.Sessions("jxc_session", store)
}

func RegisterCache() gin.HandlerFunc {
	var cacheStore persistence.CacheStore = persistence.NewRedisCache(
		config.GetEnv().RedisIP+":"+config.GetEnv().RedisPort,
		config.GetEnv().RedisPassword,
		time.Minute,
	)
	return cache.Cache(&cacheStore)
}
