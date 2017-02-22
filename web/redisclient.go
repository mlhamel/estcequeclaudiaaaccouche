package web

import (
	"gopkg.in/redis.v5"
	"net/url"
	"os"
	"time"
)

type RedisClient interface {
	Get(string) *redis.StringCmd
	Set(string, interface{}, time.Duration) *redis.StatusCmd
}

func NewClient() RedisClient {
	var value = os.Getenv("REDIS_URL")
	var password = ""

	if value == "" {
		value = "redis://localhost:6379"
	}

	u, _ := url.Parse(value)
	password, _ = u.User.Password()
	value = u.Host

	return redis.NewClient(&redis.Options{
		Addr:     value,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}
