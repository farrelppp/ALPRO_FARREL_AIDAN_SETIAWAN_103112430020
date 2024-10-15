package main

import "fmt"

func main() {
	var masukan, bil1, bil2, bil3 int

	fmt.Print("Masukan angka: ")
	fmt.Scan(&masukan)

	bil1 = masukan / 100
	bil2 = (masukan % 100) / 10
	bil3 = masukan % 10

	fmt.Println(bil1 <= bil2 && bil2 <= bil3)
}
