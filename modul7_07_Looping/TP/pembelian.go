package main

import "fmt"

func main() {
	// deklarasikan variable "barang" dan "poin" ke tipe data int
	var barang, poin int
	// menampilkan perintah untuk memasukan jumlah barang
	fmt.Print("masukan jumlah barang yang dibeli: ")
	fmt.Scan(&barang)
	// memberi nilai pada variable poin sebagai 0
	poin = 0
	// perulangan untuk menghitung poin
	for i := 1; i <= barang; i++ {
		if i <= 5 {
			poin += 10
		} else {
			poin += 15
		}
	}
	// menampilkan hasil perhitungan poin
	fmt.Printf("poin yang didapat: %d\n", poin)

}
