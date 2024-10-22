package main

import "fmt"

func main() {
	// deklarasikan variable kedalam tipe integer
	var n int
	fmt.Scan(&n)
	// mengecek apakah bilangan yang dimasukan positif atau bukan
	if n < 0 {
		// bila bukan positif maka akan muncul peringatan dibawah
		fmt.Println("harus positif")
	}
	// tetapkan nilai hasil sebagai 1
	hasil := 1
	// memulai looping untuk faktorial
	for i := 0; i < n; i++ {
		// rumus perhitungan
		hasil *= i + 1
	}
	// menampilkan hasil faktorial
	fmt.Printf("faktorial %d adalah: %d\n", n, hasil)

}
