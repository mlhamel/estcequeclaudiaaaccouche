package main

import (
  "fmt"
  "log"
  "net/http"
)

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
  addr, err := GetListenAddress()

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
