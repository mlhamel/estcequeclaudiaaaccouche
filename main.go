package main

import (
	"github.com/mlhamel/accouchement/web"
	"log"
	"net/http"
)

func main() {
	s := web.NewStatus()
	s.Refresh()

	addr, err := web.GetListenAddress()

	if err != nil {
		panic(err)
	}

	toggleStatusUrl := web.GetToggleUrl()

	http.HandleFunc("/", makeHandler(web.DisplayStatus, s))
	http.HandleFunc("/api", makeHandler(web.ApiStatus, s))
	http.HandleFunc(toggleStatusUrl, makeHandler(web.ToggleStatus, s))

	log.Printf("Listening on %s...\n", addr)
	log.Printf("Setter sets to `%s`", toggleStatusUrl)

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *web.Status), s *web.Status) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, s)
	}
}
