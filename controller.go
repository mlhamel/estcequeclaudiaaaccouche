package main

import (
	"encoding/json"
	"net/http"

	"github.com/mlhamel/accouchement/status"
)

func RenderStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	s.Refresh()

	renderTemplate(w, "templates/status.html", s.Serialize())
}

func ApiStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	json.NewEncoder(w).Encode(s.Serialize())
}

// ToggleStatus respond to an http request and toggle the status of the status manager
func ToggleStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	if s.Value() == status.No {
		s.Enable()
	} else {
		s.Disable()
	}

	json.NewEncoder(w).Encode(s.Serialize())
}

// ToggleStatusWithTwilio respond to an http request and toggle the status of the status manager
func ToggleStatusWithTwilio(w http.ResponseWriter, r *http.Request, s *status.Status) {
	if s.Value() == status.No {
		s.Enable()
	} else {
		s.Disable()
	}

	t := NewTwiML(s)

	x, err := t.Marshal()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

// RenderStatusWithTwilio respond to an http request marshal to twiml format
func RenderStatusWithTwilio(w http.ResponseWriter, r *http.Request, s *status.Status) {
	t := NewTwiML(s)
	x, err := t.Marshal()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
