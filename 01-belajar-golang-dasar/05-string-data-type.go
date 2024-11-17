package main

import "fmt"

/*
 * String is data type represent set of character
 * a string must begin with " and end with "
 */

func main() {
	fmt.Println("Muhammad")
	fmt.Println("Muhammad Fauzi")

	// To get length of a string, use len() function
	fmt.Println(len("Fauzi"))

	// To get a character from a string use index accessor [n]
	// This will print "70" a byte representation of char "70"
	// if we want to get the char represenetation we need to process first
	fmt.Println("Fauzi"[0])
}
