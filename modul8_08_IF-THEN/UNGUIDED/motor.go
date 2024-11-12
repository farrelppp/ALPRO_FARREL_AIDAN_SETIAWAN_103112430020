package main

import "fmt"

func main() {
	var orang, motor int

	fmt.Print("masukkan jumlah orang: ")
	fmt.Scan(&orang)

	motor = orang / 2
	if orang%2 != 0 {
		motor++
	}

	fmt.Printf("jumlah motor yang dibutuhkan: %d\n", motor)
}
