package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (c Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", c.Name)
}

func main() {
	cus := Customer{"Budi", "Indonesia", 20}
	cus.sayHello("Joko")
}
