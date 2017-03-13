package main

import (
	"sync"

	"github.com/mlhamel/accouchement/store"
)

const key string = "status"
const Yes string = "oui"
const No string = "non"

type StatusManager struct {
	currentValue string
	dataStore    store.Store
	mutex        *sync.Mutex
}

func NewStatusManager(dataStore store.Store, value string) *StatusManager {
	s := StatusManager{
		dataStore:    dataStore,
		currentValue: value,
		mutex:        &sync.Mutex{},
	}

	return &s
}

func (s *StatusManager) Refresh() {
	v, err := s.get(key)

	if err != nil {
		panic(err)
	} else if v == "" {
		s.currentValue = No
	} else {
		s.currentValue = v
	}
}

func (s *StatusManager) Toggle() {
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

func (s *StatusManager) Enable() {
	err := s.set(key, Yes)

	if err != nil {
		panic(err)
	}

	s.currentValue = Yes
}

func (s *StatusManager) Disable() {
	err := s.set(key, No)

	if err != nil {
		panic(err)
	}

	s.currentValue = No
}

func (s *StatusManager) Serialize() map[string]string {
	return map[string]string{"Status": s.currentValue}
}

func (s *StatusManager) Value() string {
	return s.currentValue
}

func (s *StatusManager) get(key string) (string, error) {
	s.mutex.Lock()
	value, err := s.dataStore.Get(key)
	s.mutex.Unlock()

	return value, err
}

func (s *StatusManager) set(key string, value string) error {
	s.mutex.Lock()
	err := s.dataStore.Set(key, value)
	s.mutex.Unlock()
	return err
}
