package main

import "fmt"

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, bool) {
	return "", false
}

func (s *Store) Set(key string, val any) {}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}
func main() {
	store := NewStore()
	store.Get("key1")
	store.Set("key1", "something")
	fmt.Println("GoKV - Go key value store project")
}
