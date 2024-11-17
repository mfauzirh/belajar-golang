package main

import "fmt"

/*
 * Variable is a place to store data,
 * a variable can only saved a value with same data type
 */

func main() {
	// to create a variable, use var keyword
	// var <variable_name> <data_type>
	var name string

	name = "Muhammad Fauzi"
	fmt.Println(name)

	name = "Muhammad Rizki"
	fmt.Println(name)

	// we don't need to specify data type if we assign the value right in declaration
	var age = 14
	fmt.Println(age)

	// var keyword is optional, if we need to create a variable,
	// we can simplify by using :=

	hobby := "coding"
	fmt.Println(hobby)

	hobby = "coding and watching anime"
	fmt.Println(hobby)

	// We can declare many variable in once
	var (
		firstname = "Muhammad"
		lastname  = "Fauzi"
	)

	fmt.Println(firstname, lastname)
}
