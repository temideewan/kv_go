package main

import (
	"fmt"
)

func main() {
	s := NewStore()

	s.Set("a", "42")
	s.Set("b", "72")
	a, _ := s.Get("a")
	fmt.Println(a)
	emptyKey, exists := s.Get("unknown")
	fmt.Println(emptyKey, exists)
}
