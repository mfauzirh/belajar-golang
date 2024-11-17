package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32) // carefull will resolve interger overflow

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	name := "Fauzi"
	var char = name[0]        // result of byte representation (70)
	fmt.Println(char)         // will print 70
	fmt.Println(string(char)) // we can convert the type to string will print "F" instead of 70
}
