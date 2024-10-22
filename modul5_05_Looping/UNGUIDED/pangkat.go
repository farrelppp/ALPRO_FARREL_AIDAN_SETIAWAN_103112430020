package main

import "fmt"

func main() {
	// dekalrasi kan variable kedalam tipe integer
	var a, b, hasil int

	fmt.Scan(&a, &b)
	// tetapkan nilai dari hasil
	hasil = 1
	// memulai perulangan yang dimana batasnya adalah kurang dari masukan variable b
	for i := 0; i < b; i++ {
		// rumus perhitungan
		hasil *= a

	}
	// menampilkan hasil perpangkatan
	fmt.Print("hasil perpangkatan: ", hasil)

}
