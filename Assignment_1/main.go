package main

import "fmt"

func makeRange(min int, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
func main() {
	ns := makeRange(0, 10)
	c := ""
	for _, n := range ns {
		if n%2 == 0 {
			c = "even"
		} else {
			c = "odd"
		}
		fmt.Printf("%d is %s\n", n, c)
	}
}
