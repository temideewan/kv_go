package main

import (
	"fmt"
	"sync"
	"td_redis/store/kv"
)

var (
	counter int
	mu      sync.Mutex
)

func main() {
	s := kv.NewStore(0)
	s.Set("pageviews", "0")

	const hits = 1000
	cmds := make([]Command, hits)
	for i := range cmds {
		cmds[i] = Command{Op: "INCR", Key: "pageviews"}
	}

	cmds = append(cmds, Command{Op: "INCR", Key: ""})         //empty key -> ErrEmptyKey
	cmds = append(cmds, Command{Op: "WAT", Key: "pageviews"}) //unknown op

	RestoreOnBoots(s, cmds)

	got, _ := s.Get("pageviews")
	fmt.Printf("expected: %d\n", hits)
	fmt.Printf("got: %s\n", got)
}
