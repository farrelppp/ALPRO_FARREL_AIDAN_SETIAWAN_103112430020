package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	jumlah := 0
	for i := 0; i <= n; i++ {
		genap := true
		for j := 0; j <= i; j++ {
			genap = !genap
		}
		if genap {
			jumlah++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil ", jumlah)
}
