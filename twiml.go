package main

import (
	"encoding/xml"
)

// TwiML is a Twilio status message handler
type TwiML struct {
	manager *StatusManager
}

// XML is the data structure for the TwiML manager
type XML struct {
	XMLName xml.Name `xml:"Response"`
	Say     Message  `xml:",omitempty"`
}

// Message contrain the data structure for the XML message for TwiML
type Message struct {
	String   string `xml:",innerxml"`
	Language string `xml:"language,attr"`
	Voice    string `xml:"voice,attr"`
}

// NewTwiML is creating a new TwiML
func NewTwiML(s *StatusManager) *TwiML {
	return &TwiML{manager: s}
}

// Marshal is used for serializing TwiML data
func (t *TwiML) Marshal() ([]byte, error) {
	return xml.Marshal(t.Data())
}

// Data is returning the XML representation of the values of the status manager
func (t *TwiML) Data() XML {
	value := t.manager.Value()
	message := Message{String: value, Voice: "woman", Language: "fr"}

	return XML{Say: message}
}
