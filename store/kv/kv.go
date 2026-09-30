package kv

import (
	"encoding/base64"
	"fmt"
	"maps"
	"slices"
	"td_redis/store"
	"time"
)

type Store struct {
	data    map[string]string
	maxSize int
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", store.ErrEmptyKey
	}
	val, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("Key %s does not exist", key)
	}
	return val, nil
}

func (s *Store) Len() int {
	return len(s.data)
}

func (s *Store) Set(key string, val string) error {
	if key == "" {
		return store.ErrEmptyKey
	}

	// simulate some work to slow down.
	time.Sleep(time.Second)

	_, exists := s.data[key]
	if s.maxSize > 0 && s.Len() >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q): %w", key, store.ErrStoreFull)
	}
	s.data[key] = val
	return nil
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func (s *Store) SetKeyWithEncryption(key, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := s.Set(key, encoded); err != nil {
		return "", err
	}

	return s.Get(key)

}

func NewStore(maxSize int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}

func (s *Store) Clone() *Store {
	cp := &Store{
		data:    make(map[string]string, len(s.data)),
		maxSize: s.maxSize,
	}

	cp.data = maps.Clone(s.data)
	return cp
}
