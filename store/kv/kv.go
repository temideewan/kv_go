package kv

import (
	"encoding/base64"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"sync"
	"td_redis/store"
	"time"
)

type Store struct {
	mu      sync.RWMutex
	data    map[string]string
	maxSize int
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", store.ErrEmptyKey
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	if !ok {
		return "", store.ErrKeyDoesNotExist
	}
	return val, nil
}

func (s *Store) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func (s *Store) Set(key string, val string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return store.ErrEmptyKey
	}

	// simulate some work to slow down.
	time.Sleep(time.Second)
	_, exists := s.data[key]
	if s.maxSize > 0 && len(s.data) >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q): %w", key, store.ErrStoreFull)
	}
	s.data[key] = val
	return nil
}

func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
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
	s.mu.RLock()
	defer s.mu.RUnlock()
	cp := &Store{
		data:    make(map[string]string, len(s.data)),
		maxSize: s.maxSize,
	}

	cp.data = maps.Clone(s.data)
	return cp
}

func (s *Store) Incr(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return 0, store.ErrEmptyKey
	}

	// Read
	intVal := 0
	if val, ok := s.data[key]; ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			return 0, fmt.Errorf("INCR %q: %w", key, err)
		}
		intVal = v
	}

	// Modify
	intVal++
	time.Sleep(time.Microsecond)

	// Write
	s.data[key] = strconv.Itoa(intVal)

	return intVal, nil
}
