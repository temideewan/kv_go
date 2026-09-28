package store

import "errors"

var ErrKeyDoesNotExist = errors.New("Key does not exist")
var ErrEmptyKey = errors.New("The key is mandatory")
var ErrStoreFull = errors.New("The store is currently full")

type Storer interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Len() int
	Delete(key string)
	Keys() []string
	SetKeyWithEncryption(key, value string) (string, error)
}
