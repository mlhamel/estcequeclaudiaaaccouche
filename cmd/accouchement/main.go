package main

import (
  "log"
  "net/http"

	"github.com/mlhamel/accouchement"
)

func main() {
  addr, err := accouchement.GetListenAddress()

  if err != nil {
    panic(err)
  }

	toggleStatusUrl := accouchement.GetToggleUrl()

  http.HandleFunc("/", accouchement.DisplayStatus)
  http.HandleFunc("/api", accouchement.ApiStatus)
	http.HandleFunc(toggleStatusUrl, accouchement.ToggleStatus)

  log.Printf("Listening on %s...\n", addr)
	log.Printf("Setter sets to `%s`", toggleStatusUrl)

  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
