package main

import "fmt"

func getCompleteName() (firstname, lastname string) {
	firstname = "Eko"
	lastname = "Kurniawan"

	return firstname, lastname
}

func main() {
	firstname, lastname := getCompleteName()

	fmt.Println(firstname, lastname)
}
