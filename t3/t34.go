package main

import (
	"fmt"
	"sync"
)

func main() {
	current := 0
	n := 100
	wg := new(sync.WaitGroup)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			current++
			fmt.Println(current)
			wg.Done()
		}()
		wg.Wait()
	}
}
