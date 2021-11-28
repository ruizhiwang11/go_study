package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	link := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}
	channel := make(chan string)

	for _, link := range link {
		go checkLink(link, channel)
	}

	for l := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("the website may be down")
		c <- link
		return
	}
	fmt.Println(link + " is up")

	c <- link
}
