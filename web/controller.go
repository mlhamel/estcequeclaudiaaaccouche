package web

import (
	"encoding/json"
	"net/http"
)

func DisplayStatus(w http.ResponseWriter, r *http.Request, s *Status) {
	renderTemplate(w, "templates/status.html", s.Serialize())
}

func ApiStatus(w http.ResponseWriter, r *http.Request, s *Status) {
	json.NewEncoder(w).Encode(s.Serialize())
}

func ToggleStatus(w http.ResponseWriter, r *http.Request, s *Status) {
	if s.Value() == no {
		s.Enable()
	} else {
		s.Disable()
	}

	json.NewEncoder(w).Encode(s.Serialize())
}
