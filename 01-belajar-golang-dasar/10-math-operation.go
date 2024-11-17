package main

import "fmt"

/*
 * Math operation:
 * + : addition
 * - : subtraction
 * * : multiply
 * / : division
 * % : modulo
 */

func main() {
	a := 10
	b := 10
	c := a * b
	fmt.Println(c)

	// Augmented assignment -- operation to self
	// +=, -=, *=, /=, %=

	var i = 10
	i += 10
	fmt.Println(i)

	// Unary operator
	// ++, --, -, +, !

	var j = 1
	j++
	fmt.Println(j)

	var pos = 10
	var neg = -pos

	fmt.Println(pos)
	fmt.Println(neg)
}
