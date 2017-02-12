package main

import (
  "log"
  "fmt"
  "net/url"
  "net/http"
  "os"
	"gopkg.in/redis.v5"
	"github.com/satori/go.uuid"
)

const key string = "status"
const toggleKey string = "uuid"

const yes string = "oui"
const no string = "non"

func GetRedisClient() (*redis.Client) {
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

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

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

func DisplayStatus(w http.ResponseWriter, r *http.Request) {
	status := GetStatus()

	fmt.Fprintf(w, status)
}

func ToggleStatus(w http.ResponseWriter, r *http.Request) {
	status := GetStatus()

	if status == no {
		status = EnableStatus()
	} else {
		status = DisableStatus()
	}

	fmt.Fprintf(w, status)
}

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

func main() {
  addr, err := determineListenAddress()

  if err != nil {
    panic(err)
  }

	toggleStatusUrl := GetToggleUrl()

  http.HandleFunc("/", DisplayStatus)
	http.HandleFunc(toggleStatusUrl, ToggleStatus)

  log.Printf("Listening on %s...\n", addr)
	log.Printf("Setter sets to `%s`", toggleStatusUrl)

  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
