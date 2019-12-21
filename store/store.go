package store

import "sync"

// New returns a properly initialized Store
func New() *Store {
	s := &Store{
		values: map[string]string{},
	}
	return s
}

// Store is a datastore
type Store struct {
	values map[string]string
	rwMut  sync.RWMutex
}

// Get retrives values from the store
func (s *Store) Get(key string) (string, bool) {
	s.rwMut.RLock()
	defer s.rwMut.RUnlock()

	v, ok := s.values[key]
	return v, ok
}

// Set stores values in the store
func (s *Store) Set(key, value string) {
	s.rwMut.Lock()
	defer s.rwMut.Unlock()

	s.values[key] = value
}
