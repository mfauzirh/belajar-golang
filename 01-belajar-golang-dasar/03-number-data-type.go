package main

import "fmt"

/*
 * There is 2 type of number datatype, integer and floating point
 * integer consist of:
 * 1. integer: int8, int16, int32 and int64 (negative number to positive)
 * 2. unsigned integer: uint8, uint16, uint32, and uint64 (only positive number)
 *
 * floating point consist of:
 * 1. float: float32 and float64
 * 2. complex: complex64 and complex128
 *
 * There is also an aliases for number dataype, such as:
 * 1. byte for uint8
 * 2. rune for int32
 * 3. int for minimal int32
 * 4. uint for minimal uint32
 */

func main() {
	fmt.Println("Integer number: ", 1)
	fmt.Println("Floating point number: ", 3.14)
}
