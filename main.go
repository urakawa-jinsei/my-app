package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 0

	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for range 1000 {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()

		for range 1000 {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println(n)
}
