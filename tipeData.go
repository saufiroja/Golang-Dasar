package main

import "fmt"

func main() {
	// program Hello World
	fmt.Println("Hello World")

	// Program integer atau number
	fmt.Println("satu = ", 1)
	fmt.Println("Dua koma lima = ", 2.5)

	// program Boolean
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	// Program String
	fmt.Println(len("Muhammad")) //panjang character
	fmt.Println("Muhamamd Saufi"[1])
	fmt.Println("Muhammad Saufi Roja")

	// Program declaration
	type NoKtp string
	type Married bool
	var eko NoKtp = "1233333"
	var marriedStatus Married = true
	fmt.Println(eko)
	fmt.Println(marriedStatus)
}