package main

import "fmt"

func main() {
	var berat, tinggi, hasil float64

	fmt.Print("Masukan berat badan (dalam kg): ")
	fmt.Scan(&berat)
	fmt.Print("masukan tinggi badan: ")
	fmt.Scan(&tinggi)

	hasil = berat / (tinggi * tinggi)
	fmt.Printf("Hasil kalkulasi BMI adalah %.2f kg", hasil)
}
