package main

import "fmt"


func main() {
	var (
		nama      string
		gajiPokok float64
		tunjangan float64
		potongan  float64
	)

	// print perintah untuk memasukan data karyawan
	fmt.Print("Masukkan Nama Karyawan: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan Gaji Pokok: ")
	fmt.Scan(&gajiPokok)

	fmt.Print("Masukkan Tunjangan: ")
	fmt.Scan(&tunjangan)

	fmt.Print("Masukkan Potongan: ")
	fmt.Scan(&potongan)

	// rumus hitung total gaji
	totalGaji := gajiPokok + tunjangan - potongan

	// menampilkan total gaji
	fmt.Printf("Total Gaji %s: %.2f\n", nama, totalGaji)
}
