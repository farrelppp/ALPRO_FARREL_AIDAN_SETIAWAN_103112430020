package main

import "fmt"

func main() {
	// mendeklarasikan tipe data
	var masukan, bil1, bil2, bil3 int
	// inputan user untuk ketikkan angka
	fmt.Print("Masukan angka: ")
	fmt.Scan(&masukan)
	// rumus untuk pengurutan
	bil1 = masukan / 100
	bil2 = (masukan % 100) / 10
	bil3 = masukan % 10
	// menampilkan apakah true bila urut, false bila tidak urut
	fmt.Println(bil1 < bil2 && bil2 < bil3)
}
