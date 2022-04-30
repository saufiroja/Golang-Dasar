package main

import "fmt"

func main() {
	// // kode program pass by value
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Bogor"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// // kode program pass by reference operator &
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := &address1

	// address2.City = "Bogor"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// kode program pass by reference operator *
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bogor"

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	// kode program function new
	address3 := new(Address)
	address3.City = "Bandung"
	fmt.Println(address3)
}

type Address struct {
	City, Province, Country string
}