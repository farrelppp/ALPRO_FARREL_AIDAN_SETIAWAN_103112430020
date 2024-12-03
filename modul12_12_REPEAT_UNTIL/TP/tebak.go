package main

import "fmt"

func main() {
	var tebakan int

	jawaban := 7

	for {
		fmt.Print("Masukkan angka diantara 1-10: ")
		fmt.Scan(&tebakan)

		if tebakan == jawaban {
			fmt.Print("Selamat, tebakan anda benar!")
			return
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}

}
