package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ControllerHandler func(http.ResponseWriter, *http.Request, *StatusManager)

func Serve(s *StatusManager, port string) {
	router := mux.NewRouter()

	router.HandleFunc("/", NewHandler(RenderStatus, s))
	router.HandleFunc("/api", NewHandler(APIStatus, s))

	router.
		HandleFunc("/twiml", NewHandler(RenderStatusWithTwilio, s)).
		Methods("GET")

	router.
		HandleFunc("/twiml", NewHandler(ToggleStatusWithTwilio, s)).
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

func NewHandler(fn ControllerHandler, s *StatusManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		fn(w, r, s)
	}
}
