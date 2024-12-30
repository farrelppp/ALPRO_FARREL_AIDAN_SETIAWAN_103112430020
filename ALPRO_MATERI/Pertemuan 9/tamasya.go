package main

import "fmt"

func main() {
	var N, Kapasitas, Mobil, Sisa int
	fmt.Scan(&N)

	Kapasitas = 7
	Mobil = N / Kapasitas
	Sisa = N % Kapasitas
	if N <= Kapasitas {
		fmt.Println("1 mobil dengan", Sisa, " bangku kosong")
	} else {
		fmt.Println(Mobil, " mobil penuh dan 1 Mobil berisi", Sisa, "orang")
	}
}
