package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "PZN"

	fmt.Println(book["title"])
	fmt.Println(book["author"])

	delete(book, "author")

	fmt.Println(book)
}
