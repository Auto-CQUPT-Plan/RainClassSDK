package cookieJar

import (
	"context"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

type CookieJar struct {
	miniRedisDamon *miniredis.Miniredis
	redisClient    *redis.Client
}

func (r *CookieJar) setupMiniRedisDamon() error {
	// 初始话内存数据库
	damon, err := miniredis.Run()
	if err != nil {
		return err
	}

	// 初始化客户端
	rdb := redis.NewClient(&redis.Options{Addr: damon.Addr()})

	// 测试客户端
	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		return err
	}

	r.miniRedisDamon = damon
	r.redisClient = rdb

	return nil
}

func NewCookieJar() (*CookieJar, error) {
	s := &CookieJar{}

	err := s.setupMiniRedisDamon()
	if err != nil {
		return nil, err
	}

	return s, nil
}
