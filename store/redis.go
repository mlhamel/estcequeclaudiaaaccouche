package store

import (
	"gopkg.in/redis.v5"
	"net/url"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(redisurl string, password string) *Redis {
	if redisurl == "" {
		redisurl = "redis://localhost:6379"
	}

	u, _ := url.Parse(redisurl)

	if password == "" {
		password, _ = u.User.Password()
	}

	redisurl = u.Host

	client := redis.NewClient(&redis.Options{
		Addr:     redisurl,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	return &Redis{client: client}
}

func (s *Redis) Get(key string) (string, error) {
	v, err := s.client.Get(key).Result()

	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	} else {
		return v, nil
	}
}

func (s *Redis) Set(key string, value interface{}) error {
	err := s.client.Set(key, value, 0).Err()

	if err != nil {
		return err
	}
	return nil
}
