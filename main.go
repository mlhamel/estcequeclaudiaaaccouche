package main

import (
	"github.com/mlhamel/accouchement/store"
	"github.com/mlhamel/accouchement/web"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("REDIS_URL")

	dataStore, _ := store.NewStore(store.REDIS, url, "")
	statusManager := web.NewStatus(dataStore)

	statusManager.Refresh()

	addr, err := web.GetListenAddress()

	if err != nil {
		panic(err)
	}

	toggleStatusUrl := web.GetToggleUrl()

	http.HandleFunc("/", makeHandler(web.DisplayStatus, statusManager))
	http.HandleFunc("/api", makeHandler(web.ApiStatus, statusManager))
	http.HandleFunc(toggleStatusUrl, makeHandler(web.ToggleStatus, statusManager))

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
