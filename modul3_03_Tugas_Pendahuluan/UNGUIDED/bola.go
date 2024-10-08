package main

import "fmt"

func main() {
	// deklarasi variabel jari dengan tipe data int vol, luas dengan tipe data float64
	var jari int
	var vol float64
	var luas float64

	// print perintah untuk memasukan jari jari dari lingkaran
	fmt.Print("masukan jari jari: ")
	fmt.Scan(&jari)

	// konversi jari jari ke float64
	conv := float64(jari)
	pi := 3.141592653589793
	
	// hitung volume dan luas kulit bola
	vol = (4.0 / 3.0) * pi * conv * conv * conv
	luas = 4 * pi * conv * conv

	// print hasil perhitungan volume dan luas kulit bola
	fmt.Printf("Bola dengan jejari memiliki volume %.4f dan luas kulit %.4f", vol, luas)

}
