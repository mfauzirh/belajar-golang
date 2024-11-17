package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Fauzi"
	names[2] = "Rizki"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{90, 80, 95}
	fmt.Println(values)

	// get length of an array using len()
	fmt.Println(len(values))

	var values2 = [...]int{
		90,
		85,
		90,
		10,
	}

	fmt.Println(len(values2))
}
