package status

import (
	"gopkg.in/redis.v5"
	"redisclient"
)

type Status struct {
	currentValue string
	client *redis.Client
}

const key string = "status"
const yes string = "oui"
const no string = "non"

func NewStatus() (*Status) {
	s := Status{
		client: redisclient.NewClient(),
		currentValue: no,
	}

	return &s
}

func (s *Status) Refresh() {
	v, err := s.client.Get(key).Result()

	if err == redis.Nil {
		s.currentValue = no
	} else if err != nil {
		panic(err)
	} else {
		s.currentValue = v
	}
}

func (s *Status) Enable() {
	err := s.client.Set(key, yes, 0).Err()

	if err != nil {
		panic(err)
	}

	s.currentValue = yes
}

func (s *Status) Disable() {
	err := s.client.Set(key, no, 0).Err()

	if err != nil {
		panic(err)
	}

	s.currentValue = no
}

func (s *Status) Serialize() map[string] string {
	return map[string] string {"Status": s.currentValue}
}

func (s *Status) Value() string {
	return s.currentValue
}
