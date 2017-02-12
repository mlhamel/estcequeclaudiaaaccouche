package main

import (
  "log"
  "fmt"
  "net/http"
  "os"
	"gopkg.in/redis.v5"
)

func getStatus() (string, error) {
	url := os.Getenv("REDISTOGO_URL")
	if url == "" {
		url = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client.Get("status").Result()
}

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func status(w http.ResponseWriter, r *http.Request) {
	status, err := getStatus()
	if err == redis.Nil {
		fmt.Fprintln(w, "Non")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Fprintln(w, status)
	}
}

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }

  http.HandleFunc("/", status)
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
