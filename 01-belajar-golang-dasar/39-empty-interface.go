package main

import "fmt"

// empty interface -> interface{} or have aliases any

// func Ups() interface{} {
func Ups() any {
	// return true
	// return 10
	return "ok"
}

func main() {
	fmt.Println(Ups())
}
