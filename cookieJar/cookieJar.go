package cookieJar

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

func getCookieStringKey(key string) string {
	return fmt.Sprintf("cookie:string:%s", key)
}

func getCookieValKey(key string) string {
	return fmt.Sprintf("cookie:val:%s", key)
}

func getCookieKeyKey(key string) string {
	return fmt.Sprintf("cookie:key:%s", key)
}

func (r *CookieJar) AddNewCookie(ctx context.Context, tag string, rawCookie string) error {
	if idx := strings.Index(rawCookie, ";"); idx != -1 {
		rawCookie = rawCookie[:idx]
	}

	rawCookie = strings.TrimSpace(rawCookie)
	if rawCookie == "" {
		return errors.New("empty cookie")
	}

	// 分割KV
	key, value, _ := strings.Cut(rawCookie, "=")

	tx := r.redisClient.TxPipeline()

	// 写入原始Cookie
	tx.Set(ctx, getCookieStringKey(tag), rawCookie, 0)

	// 写入Key
	tx.Set(ctx, getCookieKeyKey(tag), key, 0)

	// 写入value
	tx.Set(ctx, getCookieValKey(tag), value, 0)

	_, err := tx.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *CookieJar) GetCookieString(ctx context.Context, tag string) (string, error) {
	return r.redisClient.Get(ctx, getCookieStringKey(tag)).Result()
}

func (r *CookieJar) GetCookieKey(ctx context.Context, tag string) (string, error) {
	return r.redisClient.Get(ctx, getCookieKeyKey(tag)).Result()
}

func (r *CookieJar) GetCookieVal(ctx context.Context, tag string) (string, error) {
	return r.redisClient.Get(ctx, getCookieValKey(tag)).Result()
}
