package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var loremIpsum string
var finalIpsum string
var letterSentChan chan string

func deliverToFinal(letter string, finalIpsum *string) {
	*finalIpsum += letter
}

func capitalize(current *int, length int, letters []byte,
	finalIpsum *string, goGroup *sync.WaitGroup) {

	for *current < length {
		thisLetter := strings.ToUpper(string(letters[*current]))
		deliverToFinal(thisLetter, finalIpsum)
		*current++
	}
	goGroup.Done()
}

func main() {
	runtime.GOMAXPROCS(2)
	goGroup := new(sync.WaitGroup)
	index := new(int)
	*index = 0

	loremIpsum = "ABccc sssDDD ssdsd AAAA"

	letters := []byte(loremIpsum)
	length := len(letters)

	go capitalize(index, length, letters, &finalIpsum, goGroup)
	goGroup.Add(1)
	go func() {
		go capitalize(index, length, letters, &finalIpsum, goGroup)
		goGroup.Add(1)
	}()

	goGroup.Wait()
	fmt.Println(length, " characters.")
	fmt.Println(loremIpsum)
	fmt.Println(*index)
	fmt.Println(finalIpsum)
}
