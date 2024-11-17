package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hi, Boleh kenalan?")
	}

	// Switch juga mendukung short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// case in switch is optional

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Hello")
	case length > 5:
		fmt.Println("Hello 5")
	default:
		fmt.Println("Nama tepat")
	}
}
