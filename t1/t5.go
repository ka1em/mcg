package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func showNumber(num int, goGroup *sync.WaitGroup) {
	tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(num, tstamp)
	time.Sleep(time.Millisecond * 10)

	goGroup.Done()
}

func main() {
	iter := 10
	runtime.GOMAXPROCS(4)

	goGroup := new(sync.WaitGroup)

	for i := 0; i < iter; i++ {
		go showNumber(i, goGroup)
		goGroup.Add(1)
	}

	goGroup.Wait()
	fmt.Println("bye")
}
