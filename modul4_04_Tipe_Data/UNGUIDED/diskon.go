package main

import "fmt"

func main() {
	var diskon, total, hasil int

	fmt.Print("masukan total harga: ")
	fmt.Scan(&total)
	fmt.Print("masukan diskon: ")
	fmt.Scan(&diskon)
	hasil = total - (total * diskon / 100)

	fmt.Print("Total harga setelah diskon: ", hasil)
}
