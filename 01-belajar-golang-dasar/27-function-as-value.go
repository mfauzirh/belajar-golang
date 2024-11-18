package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	contoh := getGoodBye

	fmt.Println(contoh("Eko"))
}
