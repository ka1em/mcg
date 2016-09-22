package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"sync"
	"time"
)

var applicationStatus bool
var urls []string
var urlsProcessed int
var foundUrls []string
var fullTest string
var totalURLCount int

//var wg sync.Waitgroup

var v1 int

func readURLs(statusChannel chan int, textChannel chan string) {
	time.Sleep(time.Millisecond * 3)
	fmt.Println("Grabbing", len(urls), "urls")
	for i := 0; i < totalURLCount; i++ {
		fmt.Println("Url", i, urls[i])
		resp, _ := http.Get(urls[i])
		text, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("No HTML body")
		}
		// fmt.Println(string(text))
		textChannel <- string(text)

		statusChannel <- 0
	}
}

func addToScrapedText(textChannel chan string, processChannel chan bool) {
	for {
		select {
		case pC := <-processChannel:
			if pC == true {

			}

			if pC == false {
				close(textChannel)
				close(processChannel)
			}
		case tC := <-textChannel:
			fullTest += tC
		}
	}
}

func evaluateStatus(statusChannel chan int, textChannel chan string, processChannel chan bool) {
	for {
		select {
		case status := <-statusChannel:
			fmt.Print(urlsProcessed, totalURLCount)
			urlsProcessed++
			if status == 0 {
				fmt.Println("got url")
			}
			if status == 1 {
				close(statusChannel)
			}
			if urlsProcessed == totalURLCount {
				fmt.Println("Read all top-level URLS")
				processChannel <- false
				applicationStatus = false
			}
		}
	}
}

func main() {
	applicationStatus = true
	statusChannel := make(chan int)
	textChannel := make(chan string)
	processChannel := make(chan bool)
	totalURLCount = 0

	//wg := new(sync.WaitGroup)

	urls = append(urls, "http://www.baidu.com/index.html")
	// urls = append(urls, "http://www.mastergoco.com/index2.html")
	// urls = append(urls, "http://www.mastergoco.com/index3.html")
	// urls = append(urls, "http://www.mastergoco.com/index4.html")
	// urls = append(urls, "http://www.mastergoco.com/index5.html")

	fmt.Println("Starting spider")
	urlsProcessed = 0
	totalURLCount = len(urls)

	go evaluateStatus(statusChannel, textChannel, processChannel)
	go readURLs(statusChannel, textChannel)
	go addToScrapedText(textChannel, processChannel)

	for {
		if applicationStatus == false {
			fmt.Println(fullTest)
			fmt.Println("Done!")
			break
		}

		select {
		case sC := <-statusChannel:
			fmt.Println("Message on StatusChannel", sC)
		}
	}
}
