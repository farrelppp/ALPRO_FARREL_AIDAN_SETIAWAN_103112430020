package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi variabel 'a' dan 'b' dalam tipe float64
	var a, b float64
	var operator string
	// Perintah untuk memasukan angka
	fmt.Println("masukan angka:")
	fmt.Scanln(&a)
	// Masukan operator
	fmt.Println("masukan operator (+, -, *, /, ^):")
	fmt.Scanln(&operator)
	// Jika operator pangakt maka print teks 'masukan pangkat:'
	if operator == "^" {
		fmt.Println("masukan pangkat:")
		fmt.Scanln(&b)
		// Selain operator pangkat print teks di bawah
	} else {
		fmt.Println("masukan angka:")
		fmt.Scanln(&b)

	}
	// Switch case untuk operator yang dipilih
	switch operator {
	// case tambah
	case "+":
		fmt.Printf("Hasil penjumlahan: %.2f\n", a+b)
	// case kurang
	case "-":
		fmt.Printf("Hasil pengurangan: %.2f\n", a-b)
	// case kali
	case "*":
		fmt.Printf("Hasil perkalian: %.2f\n", a*b)
	// case bagi
	case "/":
		if b == 0 {
			fmt.Println("Error: tidak bisa bagi dengan 0")
		} else {
			fmt.Printf("Hasil pembagian: %.2f\n", a/b)
		}
	default:
		fmt.Println("Error: Invalid operator")
	// case pangkat
	case "^":
		fmt.Printf("Hasil perpangkatan: %.2f\n", math.Pow(a, b))

	}
	// Perintah untuk mengulangi program jika user ingin melakukan lagi
	var respon string
	fmt.Println("Ingin melakukan operasi lagi? (y/n)")
	fmt.Scanln(&respon)
	if respon == "y" {
		main()
	} else {
		fmt.Println("Terima kasih, sampai jumpa!")
	}
}
