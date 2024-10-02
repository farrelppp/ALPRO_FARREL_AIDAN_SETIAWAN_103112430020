package main

import "fmt"

func main() {
	var a, t float64

	fmt.Print("Masukkan alas: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&t)

	luas := 0.5 * a * t
	fmt.Printf("Luas segitiga adalah: %.2f\n", luas)

}
