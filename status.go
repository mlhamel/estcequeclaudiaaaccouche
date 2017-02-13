package accouchement

import (
	"gopkg.in/redis.v5"
)

func GetStatus() (string) {
	client := GetRedisClient()
	status, err := client.Get(key).Result()

	if err == redis.Nil {
		return no
	} else if err != nil {
		panic(err)
	} else {
		return status
	}
}

func EnableStatus() (string) {
	client := GetRedisClient()

	err := client.Set(key, yes, 0).Err()

	if err != nil {
		panic(err)
	}

	return yes
}

func DisableStatus() (string) {
	client := GetRedisClient()

	err := client.Set(key, no, 0).Err()

	if err != nil {
		panic(err)
	}

	return no
}
