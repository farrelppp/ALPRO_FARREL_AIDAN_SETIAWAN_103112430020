package main

import "fmt"

func main() {
	// deklarasikan variable "n" dan "hasil" ke tipe data int
	var n, hasil int
	fmt.Scan(&n)
	// fungsi perulangan untuk menghitung kuadrat
	for i := 1; i <= n; i++ {
		// rumusnya
		hasil = i * i
		// menampilkan hasil
		fmt.Print(hasil, " ")

	}

}
