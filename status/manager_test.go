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

	s1 := NewStatus(dataStore)
	s1.Refresh()

	if s1.Value() != No {
		t.Error("Initial value should always be no")
	}
}
