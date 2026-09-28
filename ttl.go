package main

import (
	"fmt"
	"sort"
	"time"
)

type TTLStore struct {
	data map[string]ttlEntry
}

type ttlEntry struct {
	value     string
	expiresAt time.Time
}

func NewTTlStore() *TTLStore {
	return &TTLStore{
		data: make(map[string]ttlEntry),
	}
}

func (t *TTLStore) Set(key string, value string, ttl time.Duration) error {
	if key == "" {
		return ErrEmptyKey
	}

	// insert the ttl entry, calculate the time to expire based on the ttl duration.

	return nil
}

func (t *TTLStore) GetWithTTL(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}

	entry, ok := t.data[key]
	if !ok || time.Now().After(entry.expiresAt) {
		// lazy eviction - no background job scanning for expired keys
		// deleted on demand.
		delete(t.data, key)

		return "", fmt.Errorf("key %s does not exists", key)
	}

	return entry.value, nil
}

func (t *TTLStore) Delete(key string) { delete(t.data, key) }
func (t *TTLStore) Len() int          { return len(t.data) }

func (t *TTLStore) Keys() []string {
	keys := make([]string, 0, len(t.data))

	for key := range t.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}
