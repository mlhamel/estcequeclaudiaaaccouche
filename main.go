package main

import (
  "log"
  "net/http"
	"web"
)

func main() {
  addr, err := web.GetListenAddress()

  if err != nil {
    panic(err)
  }

	toggleStatusUrl := web.GetToggleUrl()

  http.HandleFunc("/", web.DisplayStatus)
  http.HandleFunc("/api", web.ApiStatus)
	http.HandleFunc(toggleStatusUrl, web.ToggleStatus)

  log.Printf("Listening on %s...\n", addr)
	log.Printf("Setter sets to `%s`", toggleStatusUrl)

  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
