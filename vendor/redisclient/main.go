package redisclient

import (
  "net/url"
	"gopkg.in/redis.v5"
  "os"
)

func NewClient() (*redis.Client) {
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
		DB:       0,  // use default DB
	})
}
