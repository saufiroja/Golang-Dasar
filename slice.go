package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1)) // kapasitas

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "bulan") //membuat array baru jika sudah penuh array yang lama
	fmt.Println(slice3)
	slice3[1] = "bukan"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// Make slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "oja"
	newSlice[1] = "anang"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// Array vs Slice
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
