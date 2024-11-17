package main

import "fmt"

/*
 * Type declaration is way to create a new type from existing type
 * Ussualy use to make alias for existing type in purpose to make easier to understand
 */

func main() {
	type NoKTP string

	var ktpFauzi NoKTP = "32123124312324442"
	fmt.Println(ktpFauzi)
	fmt.Println(NoKTP("3244324432443244"))
}
