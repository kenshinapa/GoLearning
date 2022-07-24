package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// Golang does not support overloading (funcs with similar name even with diff. signarture)

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// // func printGreeting(sb spanishBot) {
// // 	fmt.Println(sb.getGreeting())
// // }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func (eb englishBot) getGreeting() string { // We can omit the naming of the receiver is not used
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating an spanish greeting
	return "Hola!"
}

// Concrete types vs Interface types
// You can create values directly only from concrete types.
// Interfaces are NOT generic types
// Interfaces ARE implicit
// Interfaces are a contract to help us manage types
// Interfaces are tough. Step #1 is understand how to read them
