package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	current := 0
	n := 100
	wg := new(sync.WaitGroup)
	mt := new(sync.Mutex)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			mt.Lock()
			current++
			mt.Unlock()
			fmt.Println(current)
			wg.Done()
		}()
	}
	wg.Wait()
}
