package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlhamel/accouchement/status"
)

func Serve(s *status.Status, port string) {
	router := mux.NewRouter()

	router.HandleFunc("/", makeHandler(RenderStatus, s))
	router.HandleFunc("/api", makeHandler(ApiStatus, s))

	router.
		HandleFunc("/twiml", makeHandler(RenderStatusWithTwilio, s)).
		Methods("GET")

	router.
		HandleFunc("/twiml", makeHandler(ToggleStatusWithTwilio, s)).
		Methods("Post")

	http.Handle("/", router)

	addr, err := GetListenAddress(port)

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
