package accouchement

import (
  "fmt"
	"github.com/satori/go.uuid"
	"gopkg.in/redis.v5"
)

func GenerateToggleUrl() (string) {
	u1 := uuid.NewV4()
	return fmt.Sprintf("/%s", u1)
}

func GetToggleUrl() (string) {
	client := GetRedisClient()
	toggleUrl, err := client.Get(toggleKey).Result()

	if err != redis.Nil {
		if err != nil {
			panic(err)
		}
	}

	if toggleUrl == "" {
		toggleUrl = GenerateToggleUrl()
		err = client.Set(toggleKey, toggleUrl, 0).Err()

		if err != nil {
			panic(err)
		}
	}
	return toggleUrl
}
