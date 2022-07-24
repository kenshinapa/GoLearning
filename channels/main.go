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

	c := make(chan string)

	for _, l := range links {
		// Creating a new GO routine (somewhat like a thread). We only use the go keyword in front
		// of function calls.
		go checkLink(l, c)
	}
	// This line is also a blocking call. Once any of the go routines finish, this will be executed
	// and the program will exit without waiting for the other go routines to finish.
	// fmt.Println(<-c)

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// fmt.Println(<-c) // As we only have 5 go routines, this will hang the program forever.

	// for i := 0; i < len(links); i++ {
	// Infinite loop.
	// for {
	for l := range c {
		// fmt.Println(<-c)
		// Not forever because wait for the channel to receive a message is a blocking call.
		// go checkLink(<-c, c)
		// Equical to the above line.
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down I guess"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- "Yep, it's up"
	c <- link
}

// Channels are used to communicate in-between different go routines. THE ONLY way we have to
// communicate 'em around. Channels are typed.
