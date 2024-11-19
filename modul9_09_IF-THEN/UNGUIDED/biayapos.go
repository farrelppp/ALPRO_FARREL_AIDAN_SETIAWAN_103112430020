package main

import "fmt"

func main() {
	var berat int
	var biaya int

	// perintah untuk menginput
	fmt.Print("masukkan berat (kg): ")
	fmt.Scan(&berat)

	// konversi berat ke kg dan gram
	beratkg := berat / 1000
	sisagram := berat % 1000

	// menghitung biaya kg
	biayakg := beratkg * 10000

	// menghitung biaya gram
	if beratkg > 10 {
		biaya = 0
	} else if sisagram >= 500 {
		biaya = sisagram * 5
	} else {
		biaya = sisagram * 15
	}
	totalBiaya := biayakg + biaya

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratkg, sisagram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayakg, biaya)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)

}
