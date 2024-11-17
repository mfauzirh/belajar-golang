package main

import "fmt"

// A main function should exist only one in one package,
// so multiple main func will trigger error
// For learning purpose, we'll ignore this type of error
// since we'll not gonna build this project
func main() {
	fmt.Println("This is sample text")
}
