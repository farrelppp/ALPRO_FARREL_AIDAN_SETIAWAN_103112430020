package main

import "fmt"

func main() {
	var cel, ream, fahr, kelv float64

	fmt.Print("masukan suhu dalam celcius: ")
	fmt.Scan(&cel)

	ream = cel * 4 / 5
	fahr = (cel * 9 / 5) + 32
	kelv = cel + 273.15

	fmt.Printf("Celcius: %.2f\n", cel)
	fmt.Printf("Reamur: %.2f\n", ream)
	fmt.Printf("Fahrenheit: %.2f\n", fahr)
	fmt.Printf("Kelvin: %.2f\n", kelv)
}
