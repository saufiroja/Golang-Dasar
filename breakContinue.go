package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// break (untuk menghentikan seluruh perulangan)
		if i == 5{
			break
		}
		fmt.Println(i)

		// continue (menghentikan perulangan saat ini dan akan langsung dilanjutkan ke post statement)
		if i % 2 == 0{
			continue
		}

		fmt.Println(i)
	}
}