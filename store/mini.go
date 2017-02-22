package store

import (
	"miniredis"
)

type Mini struct {
	client *Miniredis
}

func NewMini(redisurl string, password string) *Redis {
	s, _ := miniredis.Run()

	return &s
}

func (s *Redis) Get(key string) (string, error) {
	v, err := s.client.Get(key)

	return v, err
}

func (s *Redis) Set(key string, value interface{}) error {
	err := s.client.Set(key, value)

	return err
}
