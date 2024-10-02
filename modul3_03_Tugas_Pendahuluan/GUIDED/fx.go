package main

import "fmt"

func main() {
	var x, nilai int

	fmt.Print("masukan angka: ")
	fmt.Scan(&x)

	nilai = (2/(x+5) + 5)

	fmt.Println("hasil: ", nilai)
}
