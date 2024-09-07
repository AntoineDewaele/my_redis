package store

import (
	"sync"
)

type Store struct {
	data map[string]string
}

var singleton *Store
var once sync.Once

func GetStore() *Store {
	once.Do(func() {
		singleton = &Store{data: make(map[string]string)}
	})

	return singleton
}

func (s *Store) Get(key string) string {
	return s.data[key]
}

func (s *Store) Set(key string, value string) {
	s.data[key] = value
}

func (s *Store) Reset() {
	s.data = make(map[string]string)
}