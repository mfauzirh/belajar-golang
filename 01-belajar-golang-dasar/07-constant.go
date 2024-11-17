package main

import "fmt"

/*
 * Constant is a variable that can't be changed after it's initialization
 * When declaring a constant we need to specify its value
 */

func main() {
	const pi float64 = 3.14
	const pi2 = 3.14

	// pi = 3.15 will generate an error

	fmt.Println(pi)
	fmt.Println(pi2)

	const (
		firstname = "John"
		lastname  = "Doe"
	)
}
