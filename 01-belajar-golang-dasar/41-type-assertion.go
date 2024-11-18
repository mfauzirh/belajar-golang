package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()
	resultStr := result.(string)
	fmt.Println(resultStr)

	// resultInt := result.(int) // wrong type assertion can occur panic
	// fmt.Println(resultInt)

	res := random()
	switch value := res.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Other type")
	}
}
