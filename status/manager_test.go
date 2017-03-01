package status

import (
	"github.com/mlhamel/accouchement/store"
	"testing"
)

func TestNewStatus(t *testing.T) {
	t.Log("Creating new status")

	dataStore, err := store.NewStore(store.MINI, "", "")

	t.Log("Store created")

	if err != nil {
		panic(err)
	}

	s1 := NewStatus(dataStore, No)

	s1.Refresh()

	if s1.Value() != No {
		t.Error("Initial value should always be no")
	}
}

func ToggleStatus(t *testing.T) {
	t.Log("Testing Toggle status")

	dataStore, _ := store.NewStore(store.MINI, "", "")

	s1 := NewStatus(dataStore, No)

	s1.Toggle()

	if s1.Value() != Yes {
		t.Error("Toggling from no should always be yes")
	}

	s1.Toggle()

	if s1.Value() != No {
		t.Error("Toggling from yes should always be no")
	}
}

func DisableStatus(t *testing.T) {
	t.Log("Testing Disable status")

	dataStore, _ := store.NewStore(store.MINI, "", "")

	s1 := NewStatus(dataStore, Yes)

	s1.Disable()

	if s1.Value() != No {
		t.Error("Disabling from yes should always be no")
	}

	s1.Disable()

	if s1.Value() != No {
		t.Error("Disabling from no should always be no")
	}
}

func EnableStatus(t *testing.T) {
	t.Log("Testing Disable status")

	dataStore, _ := store.NewStore(store.MINI, "", "")

	s1 := NewStatus(dataStore, No)

	s1.Enable()

	if s1.Value() != Yes {
		t.Error("Enabling from no should always be yes")
	}

	s1.Enable()

	if s1.Value() != Yes {
		t.Error("Enabling from yes should always be yes")
	}
}
