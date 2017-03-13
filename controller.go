package main

import (
	"encoding/json"
	"net/http"
)

func RenderStatus(w http.ResponseWriter, r *http.Request, s *StatusManager) {
	s.Refresh()

	renderTemplate(w, "templates/status.html", s.Serialize())
}

func ApiStatus(w http.ResponseWriter, r *http.Request, s *StatusManager) {
	json.NewEncoder(w).Encode(s.Serialize())
}

// ToggleStatus respond to an http request and toggle the status of the status manager
func ToggleStatus(w http.ResponseWriter, r *http.Request, s *StatusManager) {
	if s.Value() == No {
		s.Enable()
	} else {
		s.Disable()
	}

	json.NewEncoder(w).Encode(s.Serialize())
}

// ToggleStatusWithTwilio respond to an http request and toggle the status of the status manager
func ToggleStatusWithTwilio(w http.ResponseWriter, r *http.Request, s *StatusManager) {
	if s.Value() == No {
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
func RenderStatusWithTwilio(w http.ResponseWriter, r *http.Request, s *StatusManager) {
	t := NewTwiML(s)
	x, err := t.Marshal()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
