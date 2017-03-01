package status

import (
	"github.com/mlhamel/accouchement/store"
	"sync"
)

const key string = "status"
const Yes string = "oui"
const No string = "non"

type Status struct {
	currentValue string
	dataStore    store.Store
	mutex        *sync.Mutex
}

func NewStatus(dataStore store.Store) *Status {
	s := Status{
		dataStore:    dataStore,
		currentValue: No,
		mutex:        &sync.Mutex{},
	}

	return &s
}

func (s *Status) Refresh() {
	v, err := s.get(key)

	if err != nil {
		panic(err)
	} else if v == "" {
		s.currentValue = No
	} else {
		s.currentValue = v
	}
}

func (s *Status) Toggle() {
	v, err := s.get(key)

	if err != nil {
		panic(err)
	} else if v == "" {
		s.Enable()
	} else if v == No {
		s.Enable()
	} else if v == Yes {
		s.Disable()
	}
}

func (s *Status) Enable() {
	err := s.set(key, Yes)

	if err != nil {
		panic(err)
	}

	s.currentValue = Yes
}

func (s *Status) Disable() {
	err := s.set(key, No)

	if err != nil {
		panic(err)
	}

	s.currentValue = No
}

func (s *Status) Serialize() map[string]string {
	return map[string]string{"Status": s.currentValue}
}

func (s *Status) Value() string {
	return s.currentValue
}

func (s *Status) get(key string) (string, error) {
	s.mutex.Lock()
	value, err := s.dataStore.Get(key)
	s.mutex.Unlock()

	return value, err
}

func (s *Status) set(key string, value string) error {
	s.mutex.Lock()
	err := s.dataStore.Set(key, value)
	s.mutex.Unlock()
	return err
}
