package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Split("17-08-1945", "-"))
	fmt.Println(strings.ToLower("Eko Kurniawan"))
	fmt.Println(strings.ToUpper("Eko Kurniawan"))
	fmt.Println(strings.Trim("         Eko Kurniawan     ", " "))
	fmt.Println(strings.ReplaceAll("Halo nama saya #nama# salam kenal", "#nama#", "Budi"))
}
