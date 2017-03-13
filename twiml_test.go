package main

import (
	"fmt"
	"testing"

	"github.com/mlhamel/accouchement/store"
)

func TestNewTwiml(t *testing.T) {
	dataStore, _ := store.NewStore(store.MINI, "", "")

	s1 := NewStatusManager(dataStore, No)

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
	dataStore, _ := store.NewStore(store.MINI, "", "")

	s1 := NewStatusManager(dataStore, No)

	twiml := NewTwiML(s1)
	_, err := twiml.Marshal()

	if err != nil {
		t.Error(fmt.Sprintf("Cannot marshall twiml: %s", err))
	}
}
