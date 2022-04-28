package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Eko",
		"address": "Indo",
	}

	// Menambah
	person["title"] = "programmer"

	// Mengubah
	person["name"] = "Oja"

	// Menghapus
	delete(person, "title") 

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)

	book["title"] = "Belajar Golang"
	book["author"] = "Eko"

	fmt.Println(book)
}