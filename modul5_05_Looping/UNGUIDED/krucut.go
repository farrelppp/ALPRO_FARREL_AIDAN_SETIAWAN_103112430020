package main

import "fmt"

func main() {
	// deklarasikan tipe data pada variable
	var n int

	fmt.Scan(&n)

	// looping untuk perhitungan tumus volume kerucut
	for i := 1; i <= n; i++ {
		var r, h float64
		fmt.Scan(&r, &h)
		vol := (1.0 / 3.0) * 3.14 * r * r * h
		// memunculkan hasil perhitungan dan perhitungan yang keberapa
		fmt.Printf("Luas segitiga ke-%d = %.2f\n", i, vol)
	}

}
