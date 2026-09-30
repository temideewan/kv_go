package main

import (
	"encoding/base64"
	"fmt"
	"td_redis/store"
	"td_redis/store/kv"
)

func main() {
	cmds := []Command{
		{Op: "SET", Key: "env", Value: "Production"},
		{Op: "SET", Key: "version", Value: "0.0.1"},
		{Op: "SET", Key: "debug", Value: "true"},
		{Op: "GET", Key: "env"},
		{Op: "SET", Key: "region", Value: "eu-west-1"},
		{Op: "GET", Key: "version"},
	}

	s := kv.NewStore(0)
	RestoreOnBoots(s, cmds)
	fmt.Println("restored, store size:", s.Len())
	fmt.Println("Keys:", s.Keys())
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
