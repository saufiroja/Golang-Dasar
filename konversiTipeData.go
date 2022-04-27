package main

import "fmt"

func main() {
	// konversi tiper data (1)
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// konversi tiper data (2)
	var name = "Muhammad Saufi Roja"
	var e = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}