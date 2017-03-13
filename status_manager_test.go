package main

import (
	"testing"

	"github.com/mlhamel/accouchement/store"
)

func TestNewStatusManager(t *testing.T) {
	t.Log("Creating new status")

	dataStore, err := store.NewStore(store.MINI, "", "")

	t.Log("Store created")

	if err != nil {
		panic(err)
	}

	s1 := NewStatusManager(dataStore, No)

	s1.Refresh()

	if s1.Value() != No {
		t.Error("Initial value should always be no")
	}
}

func ToggleStatusManager(t *testing.T) {
	t.Log("Testing Toggle status")

	dataStore, _ := store.NewStore(store.MINI, "", "")

	s1 := NewStatusManager(dataStore, No)

	s1.Toggle()

	if s1.Value() != Yes {
		t.Error("Toggling from no should always be yes")
	}

	s1.Toggle()

	if s1.Value() != No {
		t.Error("Toggling from yes should always be no")
	}
}

func DisableStatusManager(t *testing.T) {
	t.Log("Testing Disable status")

	dataStore, _ := store.NewStore(store.MINI, "", "")

	s1 := NewStatusManager(dataStore, Yes)

	s1.Disable()

	if s1.Value() != No {
		t.Error("Disabling from yes should always be no")
	}

	s1.Disable()

	if s1.Value() != No {
		t.Error("Disabling from no should always be no")
	}
}

func EnableStatusManager(t *testing.T) {
	t.Log("Testing Disable status")

	dataStore, _ := store.NewStore(store.MINI, "", "")

	s1 := NewStatusManager(dataStore, No)

	s1.Enable()

	if s1.Value() != Yes {
		t.Error("Enabling from no should always be yes")
	}

	s1.Enable()

	if s1.Value() != Yes {
		t.Error("Enabling from yes should always be yes")
	}
}
