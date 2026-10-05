package main

import (
	"fmt"
	"log"
	"sync"
	"td_redis/store"
)

// A parsed redis-ish operation: Op code and its arguments
type Command struct {
	Op    string
	Value string
	Key   string
}

func dispatch(s store.Storer, c Command) error {
	var err error
	switch c.Op {
	case "SET":
		log.Printf("SET key %s", c.Key)
		err = s.Set(c.Key, c.Value)
	case "Get":
		_, err = s.Get(c.Key)
	case "INCR":
		_, err = s.Incr(c.Key)
	default:
		err = fmt.Errorf("unknown op %q", c.Op)
	}
	return err
}

// This replays a slice of commands against the store
func RestoreOnBoots(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup
	errs := make(chan error, len(cmds))
	for _, c := range cmds {
		wg.Go(func() {
			errs <- dispatch(s, c)
		})
	}
	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			fmt.Printf("Error %v\n", err)
		}
	}
}
