package main

import "fmt"

func main() {
	var input int

	// print perintah untuk memasukan jumlah baris
	fmt.Print("masukan jumlah barisnya: ")
	fmt.Scan(&input)

	// cek apakah inputan bilangan bulat positif
	if input <= 0 {
		fmt.Print("inputan harus bilangan bulat positif")
		return
	}

	// rumus untuk mencetak segitiga bintang
	for i := 1; i <= input; i++ {
		for j := 1; j <= input-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		// untuk pindah baris selanjutnya
		fmt.Println()
	}
}
