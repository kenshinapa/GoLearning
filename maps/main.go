package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	// var colors map[string]string

	colors := make(map[string]string)
	// To access indivitual key we have to use the square brackets.
	// Inside the [] goes a value of correct type. colors := map[this value]string
	// No "dot" notation such as in structs
	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"

	colorsInt := make(map[int]string)
	colorsInt[0] = "#ffffff"

	fmt.Println(colors)
	fmt.Println(colorsInt)

	// Delete from map
	delete(colorsInt, 0)
	fmt.Println(colorsInt)

	// Iterate over a map
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
