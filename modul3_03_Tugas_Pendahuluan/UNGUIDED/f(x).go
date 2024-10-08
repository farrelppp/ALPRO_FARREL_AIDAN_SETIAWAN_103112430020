package main

import "fmt"

func main() {
	// dekalrasi variable di bawah dengan tipe data int
	var x int
	var nilai int
	// print perintah untuk memasukan nilai x
	fmt.Print("Masukan angka: ")
	fmt.Scan(&x)
	// perhitugan nilai x
	nilai = ((2 / (x + 5)) + 5)
	// print hasil perhitungan x
	fmt.Println("Hasil perhitungan x adalah: ", nilai)
}
