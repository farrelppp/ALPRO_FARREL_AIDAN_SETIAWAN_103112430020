package main

import "fmt"

func main() {
	var jari2, luas, keliling float64 // deklarasikan tipe data dari variabel jari2, luas, keliling

	fmt.Print("masukan jari jari lingkaran: ") // print perintah untuk memasukan jari jari lingkaran
	fmt.Scan(&jari2)

	luas = 3.14 * jari2 * jari2 // rumus luas lingkaran
	keliling = 2 * 3.14 * jari2 // rumus keliling lingkaran

	fmt.Printf("Luas lingkaran adalah %.2f dan keliling %.2f", luas, keliling) // print hasil dari perhitungan rumus luas dan keliling lingkaran
}
