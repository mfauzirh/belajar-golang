package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println(message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		fmt.Println("Ops error")
	}
}

func main() {
	runApp(true)
}
