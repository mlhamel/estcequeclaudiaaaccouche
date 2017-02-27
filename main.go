package main

import (
	"github.com/gorilla/mux"
	"github.com/mlhamel/accouchement/status"
	"github.com/mlhamel/accouchement/store"
	"github.com/mlhamel/accouchement/twilio"
	"github.com/mlhamel/accouchement/web"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("REDIS_URL")

	dataStore, _ := store.NewStore(store.REDIS, url, "")
	statusManager := status.NewStatus(dataStore)

	statusManager.Refresh()

	addr, err := web.GetListenAddress()

	if err != nil {
		panic(err)
	}

	toggleStatusUrl := web.GetToggleUrl()

	router := mux.NewRouter()

	router.HandleFunc("/", makeHandler(web.DisplayStatus, statusManager))
	router.HandleFunc("/api", makeHandler(web.ApiStatus, statusManager))
	router.HandleFunc(toggleStatusUrl, makeHandler(web.ToggleStatus, statusManager))

	router.
		HandleFunc("/twiml", makeHandler(twilio.DisplayStatus, statusManager)).
		Methods("GET")

	router.
		HandleFunc("/twiml", makeHandler(twilio.ToggleStatus, statusManager)).
		Methods("Post")

	http.Handle("/", router)

	log.Printf("Listening on %s...\n", addr)
	log.Printf("Setter sets to `%s`", toggleStatusUrl)

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *status.Status), s *status.Status) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		fn(w, r, s)
	}
}
