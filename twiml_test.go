package main

import (
	"fmt"
	"testing"
)

func TestNewTwiml(t *testing.T) {
	s1 := buildStatusManager()

	twiml := NewTwiML(s1)
	data := twiml.Data()
	message := data.Say

	if message.String != No {
		t.Error("String from message should be No")
	}

	if message.Language != "fr" {
		t.Error("Language from message should be fr")
	}

	if message.Voice != "woman" {
		t.Error("Language from message should be woman")
	}
}

func TestMarshall(t *testing.T) {
	s1 := buildStatusManager()

	twiml := NewTwiML(s1)
	_, err := twiml.Marshal()

	if err != nil {
		t.Error(fmt.Sprintf("Cannot marshall twiml: %s", err))
	}
}
