package initialize

import (
	"blog-backend/global"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.YAGAMI_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.YAGAMI_LOGGER.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		fmt.Println("[Redis]初始化成功")
		global.YAGAMI_LOGGER.Info("redis connect ping response:", zap.String("pong", pong))
		global.YAGAMI_REDIS = client
	}
}
