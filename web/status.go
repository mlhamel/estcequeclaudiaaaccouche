package web

import (
	"github.com/mlhamel/accouchement/store"
)

type Status struct {
	currentValue string
	dataStore    store.Store
}

func NewStatus(dataStore store.Store) *Status {
	s := Status{
		dataStore:    dataStore,
		currentValue: no,
	}

	return &s
}

func (s *Status) getKey(key string) (string, error) {
	return s.dataStore.Get(key)
}

func (s *Status) Refresh() {
	v, err := s.getKey(key)

	if err != nil {
		panic(err)
	} else if v == "" {
		s.currentValue = no
	} else {
		s.currentValue = v
	}
}

func (s *Status) Enable() {
	err := s.dataStore.Set(key, yes)

	if err != nil {
		panic(err)
	}

	s.currentValue = yes
}

func (s *Status) Disable() {
	err := s.dataStore.Set(key, no)

	if err != nil {
		panic(err)
	}

	s.currentValue = no
}

func (s *Status) Serialize() map[string]string {
	return map[string]string{"Status": s.currentValue}
}

func (s *Status) Value() string {
	return s.currentValue
}
