package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("firing 3 concurrent requests...")
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			makeRequests(ctx, id)
		}(i)
	}

	wg.Wait()
	fmt.Println("all done")
}

func makeRequests(ctx context.Context, id int) {
	fmt.Printf("[req %d] making the request....\n", id)
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		close(done)
	}()
	select {
	case <-done:
		fmt.Printf("[req %d] request done!\n", id)
	case <-ctx.Done():
		fmt.Printf("[req %d] request timed out! %v\n", id, ctx.Err())
	}
}
