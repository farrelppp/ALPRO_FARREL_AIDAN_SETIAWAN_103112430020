package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	if a < 0 && a%2 == 0 {
		fmt.Print("genap negatif")

	} else {
		fmt.Print("bukan genap negatif")
	}

}
