package main

import (
	"fmt"
)

func main() {
	s := NewStore(2)
	var setErr error
	setErr = s.Set("a", "42")
	setErr = s.Set("b", "72")
	setErr = s.Set("c", "24")
	a, err := s.Get("a")

	if err != nil || setErr != nil {
		fmt.Println(err)
		fmt.Println(setErr)
	}
	fmt.Println(a)
}
