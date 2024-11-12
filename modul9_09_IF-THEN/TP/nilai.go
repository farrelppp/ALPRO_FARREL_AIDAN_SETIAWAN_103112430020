package main

import "fmt"

func main() {
	var nilai int
	var huruf string

	// perintah untuk memasukan nilai siswa
	fmt.Print("masukan nilai siswa: ")
	fmt.Scan(&nilai)
	// fungsi if untuk menentukan indeks nilai siswa
	if nilai > 90 {
		huruf = "A"
	} else if nilai >= 80 && nilai <= 90 {
		huruf = "AB"
	} else if nilai >= 70 && nilai < 80 {
		huruf = "B"
	} else {
		huruf = "C"
	}
    // print hasil 
	fmt.Print("siswa mendapatkan ", huruf)
}
