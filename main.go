package main

import (
	"encoding/base64"
	"fmt"
	"td_redis/store"
	"td_redis/store/kv"
)

func main() {
	kvStore := kv.NewStore(20)
	PopulateDefaults(kvStore)
	newStore := kvStore.Clone()

	newStore.Set("env", "development")
	fmt.Println(kvStore)
	fmt.Println(newStore)

}

func PopulateDefaults(s store.Storer) error {
	defaults := map[string]string{
		"env":     "production",
		"version": "0.0.1",
		"debug":   "true",
	}

	for k, v := range defaults {
		if err := s.Set(k, v); err != nil {
			return fmt.Errorf("PopulateDefaults %w", err)
		}
	}
	return nil
}

func CreateStore() store.Storer {
	plain := kv.NewStore(20)
	logger := NewLoggingMiddleware(plain)
	metric := NewMetricMiddleware(logger)
	return metric
}

func SetKeyWithEncryption(s store.Storer, key, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := s.Set(key, encoded); err != nil {
		return "", err
	}

	return s.Get(key)

}
