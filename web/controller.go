package web

import (
	"encoding/json"
	"github.com/mlhamel/accouchement/status"
	"net/http"
)

func DisplayStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	s.Refresh()

	renderTemplate(w, "templates/status.html", s.Serialize())
}

func ApiStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	json.NewEncoder(w).Encode(s.Serialize())
}

func ToggleStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	if s.Value() == status.No {
		s.Enable()
	} else {
		s.Disable()
	}

	json.NewEncoder(w).Encode(s.Serialize())
}
