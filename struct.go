package main

import "fmt"

func main() {
	// Program data struct
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indo"
	eko.Age = 20
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	// program struct literals
	budi := Customer {
		Name: "Budi",
		Address: "Indo",
		Age: 20,
	}
	fmt.Println(budi)

	// struct method
	oja := Customer{
		Name: "Oja",
	}
	oja.sayHalo("eko")
}

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHalo(name string) {
	fmt.Println("Hello,", name ,"My name is", customer.Name)
}