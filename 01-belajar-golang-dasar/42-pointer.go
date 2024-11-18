package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	addr1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	addr2 := &addr1

	addr2.City = "Jakarta"

	// addr2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*addr2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(addr1)
	fmt.Println(addr2)
}
