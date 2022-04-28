package main

import "fmt"

func main() {
	var name = "budi"

	if name == "eko" {
		fmt.Println("Hello eko")
	} else if name == "budi"{
		fmt.Println("Hello budi")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	// short statement
	if length := len(name); length > 5{
		fmt.Println("Name terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}