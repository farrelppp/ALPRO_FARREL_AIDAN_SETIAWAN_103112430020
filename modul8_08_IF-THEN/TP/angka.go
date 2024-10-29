package main

import "fmt"

func main() {
	var a int
	// memunculkan perintah untuk memasukan angka
	fmt.Print("masukan angka: ")
	fmt.Scan(&a)
	// mengecek apakah angka genap atau ganjil dan menampilkan hasilnya
	if a%2 == 0 {
		fmt.Printf("%d adalah bilangan genap.\n", a)
	} else {
		fmt.Printf("%d adalah bilangan ganjil.\n", a)
	}
}
