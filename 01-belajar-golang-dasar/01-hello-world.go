// A main function in go must go inside package called "main"
package main

// To print a text into terminal, use package "fmt"
import "fmt"

// function "main" serve as entry point for your application,
// there should be one main function for your app
func main() {
	// To print a text into console with new line using fmt,
	// use "Println" function and pass the text that want to print as an argument
	fmt.Println("Hello World")
}
