package main

import "fmt"

func main() {
	var masukan string

	const benar = "password123"

	const Percobaan = 3

	percobaan := 0

	for percobaan < Percobaan {
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&masukan)

		if masukan == benar {
			fmt.Println("Login berhasil!")
			return
		}
		percobaan++
		fmt.Printf("Password salah. Sisa percobaan: %d\n", Percobaan-percobaan)
	}

	fmt.Println("Gagal")
}
