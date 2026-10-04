package main

import (
	"encoding/base64"
	"fmt"
	"sync"
	"td_redis/store"
	"td_redis/store/kv"
	"time"
)

var (
	counter int
	mu      sync.Mutex
)

func increment() {
	mu.Lock()
	defer mu.Unlock()
	v := counter
	time.Sleep(time.Microsecond)
	counter = v + 1
}

func main() {
	expected := 1000
	var wg sync.WaitGroup

	for i := 0; i < expected; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter)
	fmt.Println("expected:", expected)
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
