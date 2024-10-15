package main

import "fmt"

func main() {
	var bmi, tinggi float64

	fmt.Print("masukan bmi: ")
	fmt.Scan(&bmi)
	fmt.Print("masukan tinggi: ")
	fmt.Scan(&tinggi)

	hasil := bmi * (tinggi * tinggi)

	fmt.Print("berat badannya adalah: kg\n", hasil)

}
