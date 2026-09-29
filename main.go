package main

import (
	"encoding/base64"
	"fmt"
	"td_redis/store"
	"td_redis/store/kv"
	"td_redis/store/ttl"
	"time"
)

func main() {
	keyValueStore := kv.NewStore(20)

	encrypted, err := SetKeyWithEncryption(keyValueStore, "a", "This is an encrypted text")
	if err != nil {
		fmt.Println(err)
		return
	}

	ttlStore := ttl.NewTTLStore(time.Second * 2)

	encryptedTTl, err := SetKeyWithEncryption(ttlStore, "a", "This is a n encrypted text")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encrypted)
	fmt.Println(encryptedTTl)
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
