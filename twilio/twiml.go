package twilio

import (
	"encoding/xml"
	"github.com/mlhamel/accouchement/status"
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

func NewTwiML(s *status.Status) *TwiML {
	message := Message{String: s.Value(), Voice: "woman", Language: "fr"}
	return &TwiML{Say: message}
}

func (t *TwiML) Marshal() ([]byte, error) {

	return xml.Marshal(t)
}
