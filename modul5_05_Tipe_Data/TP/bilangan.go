package main

import "fmt"

func main() {
	var input, jum int

	// print perintah input untuk pengguna
	fmt.Print("masukan bilangan bulat positif: ")
	fmt.Scan(&input)

	// cek apakah bilangan yang dimasukan bulat positif

	if input <= 0 {
		fmt.Println("harus bilangan bulat positif.")
		return
	}

	// rumus jumlah deret
	for i := 1; i <= input; i++ {
		jum += i
	}

	// menampilkan hasil
	fmt.Printf("Jumlah deret angka 1 sampai %d adalah %d ", input, jum)

}
