package store

import (
	"github.com/alicebob/miniredis"
)

type Mini struct {
	client *miniredis.Miniredis
}

func NewMini() *Mini {
	client := miniredis.NewMiniRedis()

	err := client.Start()

	if err != nil {
		panic(err)
	}

	return &Mini{client: client}
}

func (s *Mini) Get(key string) (string, error) {
	v, err := s.client.Get(key)

	if err == miniredis.ErrKeyNotFound {
		return "", nil
	} else if err != nil {
		return "", err
	} else {
		return v, nil
	}
}

func (s *Mini) Set(key string, value interface{}) error {
	err := s.client.Set(key, value.(string))

	return err
}
