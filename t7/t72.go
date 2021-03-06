package main

import (
	"fmt"
)

func abs(fxChan chan func() string) {
	fxChan <- func() string {
		return "Sent!"
	}
}

func main() {
	fxChan := make(chan func() string)
	defer close(fxChan)
	go abs(fxChan)
	select {
	case rfx := <-fxChan:
		msg := rfx()
		fmt.Println(msg)
		fmt.Println("Received")
	}
}
