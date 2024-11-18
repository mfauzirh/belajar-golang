package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal" // blank identifier, so if we want only trigger init method without using other method, so go not complaint
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
