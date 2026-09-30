package main

import (
	"log"
	"td_redis/store"
)

// A parsed redis-ish operation: Op code and its arguments
type Command struct {
	Op    string
	Value string
	Key   string
}

func dispatch(s store.Storer, c Command) {
	switch c.Op {
	case "SET":
		log.Printf("SET key %s", c.Key)
		s.Set(c.Key, c.Value)
	case "Get":
		s.Get(c.Key)
	}
}

// This replays a slice of commands against the store
func RestoreOnBoots(s store.Storer, cmds []Command) {
	for _, c := range cmds {
		dispatch(s, c)
	}
}
