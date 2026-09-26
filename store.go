package main

import (
	"errors"
	"fmt"
	"slices"
)

var ErrKeyDoesNotExist = errors.New("Key does not exist")
var ErrEmptyKey = errors.New("The key is mandatory")
var ErrMaxLengthExceeded = errors.New("The max length is exceeded")

type Store struct {
	data    map[string]string
	maxSize int
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
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
		return ErrEmptyKey
	}

	_, exists := s.data[key]
	if s.maxSize > 0 && s.Len() >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q): %w", key, ErrMaxLengthExceeded)
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

func NewStore(maxSize int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}
