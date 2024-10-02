package main

import "fmt"

func main() {
	var vol int

	fmt.Print("masukan volume: ")
	fmt.Scan(&vol)

	hasil := vol * vol * vol
	println("hasil: ", hasil)
}
