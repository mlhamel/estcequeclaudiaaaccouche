package main

import (
	"github.com/gorilla/mux"
	"github.com/mlhamel/accouchement/status"
	"github.com/mlhamel/accouchement/twilio"
	"github.com/mlhamel/accouchement/web"
	"log"
	"net/http"
)

func Serve(s *status.Status, port string) {
	router := mux.NewRouter()

	router.HandleFunc("/", makeHandler(web.DisplayStatus, s))
	router.HandleFunc("/api", makeHandler(web.ApiStatus, s))

	router.
		HandleFunc("/twiml", makeHandler(twilio.DisplayStatus, s)).
		Methods("GET")

	router.
		HandleFunc("/twiml", makeHandler(twilio.ToggleStatus, s)).
		Methods("Post")

	http.Handle("/", router)

	addr, err := web.GetListenAddress(port)

	if err != nil {
		panic(err)
	}

	log.Printf("Listening on %s...\n", addr)

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
