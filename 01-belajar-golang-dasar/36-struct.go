package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var eko Customer
	eko.Name = "Eko Kurniawan"
	eko.Address = "Jln. Sucipta"
	eko.Age = 30

	fmt.Println(eko)
	fmt.Println(eko.Name)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 28}
	fmt.Println(budi)
}
