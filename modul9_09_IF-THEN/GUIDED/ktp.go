package main

import "fmt"

func main() {
	var usia int
	var kk bool

	fmt.Scan(&usia)
	fmt.Scan(&kk)

	if usia >= 17 && kk == true {
		fmt.Println("Anda dapat membuat ktp")
	} else {
		fmt.Println("Anda tidak dapat membuat ktp")
	}
}
