package main

import "fmt"

func main() {
	var b int
	var prima bool

	// perintah untuk menginput bilangan bulat
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	if b <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1.")
		return
	}

	// rumus untuk cari faktor bilangan
	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	// menemtukan bilangan prima
	prima = true
	for i := 2; i < b; i++ {
		if b%i == 0 {
			prima = false
			break
		}
	}
	if b == 1 {
		prima = false
	}

	// mengecek dan menampilkan hasil
	if prima {
		fmt.Println("Prima: ", true)
	} else {
		fmt.Println("Prima: ", false)
	}
}
