package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64

	// Input dari pengguna
	fmt.Println("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Konversi dari Fahrenheit ke Celsius
	celsius := (fahrenheit - 32) * 5 / 9

	// Menampilkan hasil
	fmt.Printf("Suhu %.0fF dalam Celsius adalah %.0fC\n", fahrenheit, celsius)
}
