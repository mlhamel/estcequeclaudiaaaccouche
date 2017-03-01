package twilio

import (
	"github.com/mlhamel/accouchement/status"
	"net/http"
)

func ToggleStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
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

func DisplayStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	t := NewTwiML(s)
	x, err := t.Marshal()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
