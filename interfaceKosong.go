package main

import "fmt"

func main() {
	kosong := Ups(3)
	fmt.Println(kosong)
}

func Ups(i int) interface{} { // interface {} bisa di ganti dengan any
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	} else {
		return "Ups"
	}
}