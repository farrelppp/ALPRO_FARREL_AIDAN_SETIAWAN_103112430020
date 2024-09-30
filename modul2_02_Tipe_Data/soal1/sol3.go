package main

import "fmt"

func main() {
	var jari2 float64

	fmt.Println("masukan jari jari lingkaran =")
	fmt.Scanln(&jari2)
	luas := 3.14 * jari2 * jari2
	fmt.Printf("Luas lingkaran dengan jari-jari %f adalah %.2f \n", jari2, luas)

}
