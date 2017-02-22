package store

import (
	"errors"
)

const (
	REDIS = iota
	MINI
)

type Store interface {
	Get(string) (string, error)
	Set(string, interface{}) error
}

func NewStore(t int, url string, password string) (Store, error) {
	switch t {
	case REDIS:
		return NewRedis(url, password), nil
	default:
		return nil, errors.New("Invalid Store Type")
	}
}
