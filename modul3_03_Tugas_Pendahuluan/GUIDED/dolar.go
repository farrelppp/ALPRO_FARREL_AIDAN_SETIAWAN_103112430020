package main

import "fmt"

func main() {
	var idr int

	fmt.Print("masukan rupiah: ")
	fmt.Scan(&idr)

	hasil := idr / 15000
	println("dolar: ", hasil, "$")
}
