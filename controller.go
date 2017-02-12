package main

import (
  "fmt"
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
