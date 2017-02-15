package web

import (
  "fmt"
	"github.com/satori/go.uuid"
	"gopkg.in/redis.v5"
)

func GenerateToggleUrl() (string) {
	u1 := uuid.NewV4()
	return fmt.Sprintf("/%s", u1)
}

func SetToggleUrl(url string) () {
	client := GetRedisClient()
	err := client.Set(toggleKey, url, 0).Err()

	if err != nil {
		panic(err)
	}
}

func GetToggleUrl() (string) {
	client := GetRedisClient()
	url, err := client.Get(toggleKey).Result()

	if err != redis.Nil && err != nil {
		panic(err)
	}

	if url == "" {
		url = GenerateToggleUrl()
		SetToggleUrl(url)
	}

	return url
}
