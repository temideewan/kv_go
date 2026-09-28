package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	keyValueStore := NewStore(20)

	encrypted, err := SetKeyWithEncryption(*keyValueStore, "a", "This is an encrypted text")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encrypted)
}

func SetKeyWithEncryption(s Store, key, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := s.Set(key, encoded); err != nil {
		return "", err
	}

	return s.Get(key)

}
