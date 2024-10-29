package main

import "fmt"

func main() {
	var a int

	// memunculkan perintah untuk memasukan nilai
	fmt.Print("masukan nilai: ")
	fmt.Scanln(&a)
	// mengecek apakah nilai lebih besar sama dengan 70 atau lebih kecil dari 70
	if a >= 70 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}
