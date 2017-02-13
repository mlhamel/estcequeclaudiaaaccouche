package main

import (
  "log"
  "net/http"
)

func main() {
  addr, err := GetListenAddress()

  if err != nil {
    panic(err)
  }

	toggleStatusUrl := GetToggleUrl()

  http.HandleFunc("/", DisplayStatus)
  http.HandleFunc("/api", ApiStatus)
	http.HandleFunc(toggleStatusUrl, ToggleStatus)

  log.Printf("Listening on %s...\n", addr)
	log.Printf("Setter sets to `%s`", toggleStatusUrl)

  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
