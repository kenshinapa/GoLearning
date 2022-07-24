package main

import (
	"fmt"
)

// Defining structs. We do not use commas
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	// This other way is also valid
	contactInfo
}

func main() {
	// Here, we use commas. On the definition, we don't
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// contact: contactInfo{
		// 	email:   "jim@gmail.com",
		// 	zipCode: 77777,
		// },
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 77777,
		},
	}
	// jimPointer := &jim
	// jimPointer.updateName("Arty")
	// (&jim).updateName("Arty")
	jim.updateName("Arty")
	jim.print()

	// Value types are int, float, string, bool, structs
	// There are reference types. (slices, maps, channels, pointers, functions)
	myStrSlice := []string{"hello", "there"}
	updateFstStrSliceElement(myStrSlice)
	fmt.Println(myStrSlice)

	name := "Bob"
	namePtr := &name
	fmt.Println(&name)
	fmt.Println(namePtr)
	fmt.Println(&namePtr)
	printPointer(namePtr)
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateFstStrSliceElement(s []string) {
	s[0] = "Bye"
}

func printPointer(strPtr *string) {
	fmt.Println(&strPtr)
}
