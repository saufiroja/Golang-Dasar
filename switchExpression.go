package main

import "fmt"

func main() {
	name := "budi"

	switch name {
	case "eko":
		fmt.Println("Hello eko")
	case "budi":
		fmt.Println("Hello budi")
	default:
		fmt.Println("Hi, Boleh kenalan dulu?")
	}

	// short statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// } 

	length := len(name)
	switch{
	case length > 10:
		fmt.Println("Nama sangat panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}