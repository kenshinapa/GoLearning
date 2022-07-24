package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com.mx",
	}
	for _, l := range links {
		// Creating a new GO routine (somewhat like a thread). We only use the go keyword in front
		// of function calls.
		go checkLink(l)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

// Channels are used to communicate in-between different go routines. THE ONLY way we have to
// communicate 'em around. Channels are typed.
