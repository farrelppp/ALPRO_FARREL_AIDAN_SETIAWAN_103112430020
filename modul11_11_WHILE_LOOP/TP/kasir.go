package main

import "fmt"

func main() {
	var namaItem string
	var hargaItem float64
	var jumlah float64
	var lanjut string

	for {
		fmt.Print("Masukkan nama item: ")
		fmt.Scanln(&namaItem)
		fmt.Print("Masukkan harga item: ")
		fmt.Scanln(&hargaItem)

		jumlah += hargaItem

		fmt.Print("Lanjut belanja? (y/n): ")
		fmt.Scanln(&lanjut)
		if lanjut != "y" {
			break
		}
	}
	fmt.Printf("Total harga belanjaan: %.2f\n", jumlah)
	fmt.Printf("Terima kasih sudah belanja")
}
