package main

import "fmt"

func main() {
	var berat1, berat2 float64

	for {

		fmt.Println("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&berat1, &berat2)

		if berat1 < 0 || berat2 < 0 || (berat1+berat2) > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := berat1 - berat2
		if selisih < 0 {
			selisih = -selisih
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
