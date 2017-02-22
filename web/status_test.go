package web

import (
	"testing"
)

func TestNewStatus(t *testing.T) {
	t.Log("Creating new status")

	dataStore, _ := store.NewStore(store.MINI, url, "")

	if err != nil {
		panic(err)
	}

	s1 := NewStatus(dataStore)
	s1.Refresh()

	if s1.Value() != no {
		t.Error("Initial value should always be no")
	}
}
