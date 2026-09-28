package main

import (
	"fmt"
)

func main() {
	s := NewStore(2)
	_ = s.Set("a", "42")
	_ = s.Set("b", "72")
	if setErr := s.Set("c", "24"); setErr != nil {
		fmt.Println(setErr)
	}
	a, err := s.Get("a")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
