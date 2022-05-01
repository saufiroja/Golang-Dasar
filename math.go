package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40)) //di bulatkan ke atas
	fmt.Println(math.Floor(1.60)) //di bulatkan ke bawah
	fmt.Println(math.Round(1.5)) // di bulatkan misal 1 - 1.4 akan kebawah jika 1.5 - 2 akan keatas
	fmt.Println(math.Max(1,2))
	fmt.Println(math.Min(1,2))
}