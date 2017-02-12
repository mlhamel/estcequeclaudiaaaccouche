package main

import (
  "log"
  "fmt"
  "net/http"
  "os"
	"gopkg.in/redis.v5"
	"github.com/satori/go.uuid"
)

const key string = "status"
const yes string = "oui"
const no string = "non"

func GetRedisClient() (*redis.Client) {
	url := os.Getenv("REDIS_URL")
	if url == "" {
		url = "localhost:6379"
	}
	return redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func getStatus() (string, error) {
	client := GetRedisClient()

	return client.Get(key).Result()
}

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func GetStatus() (string) {
	status, err := getStatus()
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

func main() {
  addr, err := determineListenAddress()

  if err != nil {
    panic(err)
  }

	u1 := uuid.NewV4()
	toggleStatusUrl := fmt.Sprintf("/%s", u1)

  http.HandleFunc("/", DisplayStatus)
	http.HandleFunc(toggleStatusUrl, ToggleStatus)

  log.Printf("Listening on %s...\n", addr)
	log.Printf("Setter sets to `%s`", toggleStatusUrl)

  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
