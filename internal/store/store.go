package store

import (
	"sync"
	"time"
)

type StoreData struct {
	data string
	ttl int64
}

type Store struct {
	data map[string]StoreData
}

var singleton *Store
var once sync.Once

func GetStore() *Store {
	once.Do(func() {
		singleton = &Store{data: make(map[string]StoreData)}
	})

	return singleton
}

func (s *Store) Get(key string) string {
	s.deleteIfExpired(key)

	return s.data[key].data
}

func (s *Store) GetTTL(key string) int64 {
	s.deleteIfExpired(key)

	return s.data[key].ttl
}

func (s *Store) GetData() map[string]StoreData {
	return s.data
}

func (s *Store) Set(key string, value string) {
	s.data[key] = StoreData{data: value}
}

func (s *Store) SetWithTTL(key string, value string, ttl int64) {
	s.data[key] = StoreData{data: value, ttl: ttl}
}

func (s *Store) Reset() {
	s.data = make(map[string]StoreData)
}

func (s *Store) deleteIfExpired(key string) {
	if s.data[key].ttl > 0 && time.Now().Unix() > s.data[key].ttl {
		delete(s.data, key)
	}
}