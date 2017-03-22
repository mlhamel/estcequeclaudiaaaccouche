package main

import (
	"sync"

	"github.com/mlhamel/accouchement/store"
)

const key string = "status"
const urlKey string = "statusImage"
const Yes string = "oui"
const No string = "non"

type StatusManager struct {
	currentValue     string
	imageURL         string
	authorizedSource string
	dataStore        store.Store
	mutex            *sync.Mutex
}

type SerializedStatus map[string]string

func NewStatusManager(dataStore store.Store, value string, source string) *StatusManager {
	s := StatusManager{
		dataStore:        dataStore,
		currentValue:     value,
		imageURL:         "",
		authorizedSource: source,
		mutex:            &sync.Mutex{},
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

func (s *StatusManager) SetImage(url string) {
	err := s.set(urlKey, No)

	if err != nil {
		panic(err)
	}

	s.imageURL = url
}

func (s *StatusManager) GetAuthorization(source string) bool {
	if s.authorizedSource == "" {
		return true
	} else {
		return source == s.authorizedSource
	}
}

func (s *StatusManager) Serialize() SerializedStatus {
	return SerializedStatus{"Status": s.currentValue, "Image": s.imageURL}
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
