package main

import "fmt"

func main() {
	var usia int
	// menampilkan perintah untuk memasukan usia
	fmt.Print("Masukkan usia anda: ")
	fmt.Scan(&usia)
	// switch case untuk menentukan kategori anda berdasarkan masukan umur
	switch {
	case usia >= 0 && usia <= 12:
		fmt.Print("Kategori: Anak-Anak")
	case usia >= 13 && usia <= 17:
		fmt.Print("Kategori: Remaja")
	case usia >= 18 && usia <= 64:
		fmt.Print("Kategori: Dewasa")
	case usia >= 65:
		fmt.Print("Kategori: Lansia")
	default:
		fmt.Print("Usia tidak valid")
	}
}
