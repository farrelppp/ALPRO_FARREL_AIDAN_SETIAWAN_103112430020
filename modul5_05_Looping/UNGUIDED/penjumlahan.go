package main

import "fmt"

func main() {
	// deklarasikan tipe data pada variable
	var n, hasil int

	fmt.Scan(&n)

	// looping untuk penjumlahan
	for i := 1; i <= n; i++ {
		hasil += i
	}
	// menampilkan hasil
	fmt.Print("hasil penjumlahan: ", hasil)
}
