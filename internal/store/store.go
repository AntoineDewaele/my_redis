package store

type Store struct {
	data map[string]string
}

func New() *Store {
	return &Store{data: make(map[string]string)}
}

func (s *Store) Get(key string) string {
	return s.data[key]
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}