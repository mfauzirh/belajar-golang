package main

import "fmt"

func getFullName() (string, string) {
	return "Eko", "Kurniawan"
}

func main() {
	firstname, lastname := getFullName()
	fmt.Println(firstname, lastname)

	// ignore return value using _
	firstname2, _ := getFullName()
	fmt.Println(firstname2)
}
