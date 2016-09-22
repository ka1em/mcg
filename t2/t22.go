package main

import (
	"fmt"
	"strings"
)

func shortenString(message string) func() string {
	return func() string {
		messageSlice := strings.Split(message, " ")
		worldLength := len(messageSlice)
		if worldLength < 1 {
			return "Nothing Left"
		} else {
			messageSlice = messageSlice[:(worldLength - 1)]
			message = strings.Join(messageSlice, " ")
			return message
		}
	}
}

func main() {
	myString := shortenString("Welcome sdssd sddssds sds sds sds sdsd sds")

	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())

}
