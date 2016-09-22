package main

import (
	"fmt"
	"strings"
	//	"sync"
)

var initialString string
var initialBytes []byte
var stringLength int
var finalString string
var lettersProcessed int
var applicationStatus bool

//var wg sync.WaitGroup

func getLetters(gQ chan string) {
	for i := range initialBytes {
		gQ <- string(initialBytes[i])
	}
}

func capitalizeLetters(gQ chan string, sQ chan string) {
	for {
		if lettersProcessed >= stringLength {
			applicationStatus = false
			break
		}

		select {
		case letter := <-gQ:
			capitalLetter := strings.ToUpper(letter)
			finalString += capitalLetter
			lettersProcessed++
		}
	}
}

func main() {
	applicationStatus = true

	getQueue := make(chan string)
	stackQueue := make(chan string)

	initialString = "AA xxxx asdsdsdsd  ss"
	initialBytes = []byte(initialString)
	stringLength = len(initialString)
	lettersProcessed = 0

	fmt.Println("start")

	go getLetters(getQueue)
	capitalizeLetters(getQueue, stackQueue)

	close(getQueue)
	close(stackQueue)

	for {
		if applicationStatus == false {
			fmt.Println("Done")
			fmt.Println(finalString)
			break
		}
	}
}
