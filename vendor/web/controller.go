package web

import (
	"encoding/json"
  "net/http"
	"status"
)

func DisplayStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	renderTemplate(w, "templates/status.html", s.Serialize())
}

func ApiStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	json.NewEncoder(w).Encode(s.Serialize())
}

func ToggleStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	if s.Value() == no {
		s.Enable()
	} else {
		s.Disable()
	}

	json.NewEncoder(w).Encode(s.Serialize())
}
