package main

import "fmt"

func main() {
	// Deklarasi variabel 'fahrenheit' dalam tipe float64
	var fahrenheit float64
	// Print teks untuk masukan suhu fahrenheit
	fmt.Println("masukan suhu fahrenheit:")
	fmt.Scanln(&fahrenheit)
	// Mengkonversi suhu fahrenheit ke kelvin
	kelvin := (fahrenheit-32)*5/9 + 273.15
	// Menampilkan hasil konversi
	fmt.Printf("Hasil konversi dari fahrenheit ke kelvin adalah: %.2f K\n", kelvin)
}
