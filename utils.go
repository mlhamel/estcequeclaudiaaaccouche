package main

import (
	"html/template"
	"net/http"

	"github.com/mlhamel/accouchement/store"
)

func GetListenAddress(port string) (string, error) {
	return ":" + port, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		t.Execute(w, p)
	}
}

func buildStatusManager() *StatusManager {
	dataStore, _ := store.NewStore(store.MINI, "", "")
	statusManager := NewStatusManager(dataStore, No, "")

	return statusManager
}
