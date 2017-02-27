package twilio

import (
	"encoding/xml"
	"github.com/mlhamel/accouchement/status"
	"net/http"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Say     Message  `xml:",omitempty"`
}

type Message struct {
	String   string `xml:",innerxml"`
	Language string `xml:"language,attr"`
	Voice    string `xml:"voice,attr"`
}

func getXmlStatus(s *status.Status) ([]byte, error) {
	message := Message{String: s.Value(), Voice: "woman", Language: "fr"}
	twiml := TwiML{Say: message}

	return xml.Marshal(twiml)
}

func ToggleStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	if s.Value() == status.No {
		s.Enable()
	} else {
		s.Disable()
	}

	x, err := getXmlStatus(s)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

func DisplayStatus(w http.ResponseWriter, r *http.Request, s *status.Status) {
	x, err := getXmlStatus(s)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
