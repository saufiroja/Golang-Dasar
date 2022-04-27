package main

import "fmt"

func main(){
	var nilai = 90
	var absensi = 80

	var nilaiAkhir bool = nilai > 80
	var absensiAkhir bool = absensi > 80

	var lulus bool = nilaiAkhir && absensiAkhir

	fmt.Println(lulus)
}