package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// for <init_statement>; <condition>; <post_statement>
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// for range
	names := []string{"Eko", "Kurniawan", "Khannedy"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
