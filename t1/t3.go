package main

//import "os"
import "fmt"

func main() {
	/*
		file, _ := os.Create("defer.text")

		defer file.Close()

		for {
			break
		}
	*/

	aValue := new(int)

	//return 0 not 100, it is the default
	//value for an integer
	defer fmt.Println(*aValue)

	for i := 0; i < 100; i++ {
		*aValue++
	}
}
