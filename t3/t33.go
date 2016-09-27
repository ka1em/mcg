package main

import (
	"fmt"
)

type intInterface struct {
}

type stringInterface struct {
}

func (number intInterface) Add(a, b int) int {
	return a + b
}

func (str stringInterface) Add(a, b string) string {
	return a + b
}

func main() {
	number := new(intInterface)
	x := number.Add(1, 2)
	fmt.Println(x)

	str := new(stringInterface)
	s := str.Add("www", "xxx")
	fmt.Println(s)
}
