package main

import "fmt"

func main() {
	//Deklarasi variabel 'psisi' dalam tipe float64 dan set nilai psisi 27
	var psisi float64 = 27
	//Menghitung dengan menggunakan rumus keliling dan luas persegi
	keliling := 4 * psisi
	luas := psisi * psisi
	//Menampilkan hasil
	fmt.Printf("Hasil dari perhitungan menggunakan rumus keliling persegi untuk alun alun purwokerto adalah: %.2f m\n", keliling)
	fmt.Printf("Hasil dari perhitungan menggunakan rumus luas persegi untuk alun alun purwokerto adalahh: %.2f m\n", luas)
}
