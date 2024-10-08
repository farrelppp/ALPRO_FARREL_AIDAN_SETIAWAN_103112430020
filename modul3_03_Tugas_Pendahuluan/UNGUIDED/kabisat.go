package main

import "fmt"

func main() {
	// deklarasikan variable tahun dengan tipe data int
	var tahun int
	// print perintah untuk memasukan tahun
	fmt.Print("Masukan tahun: ")
	fmt.Scan(&tahun)
	// rumus untuk mengecek apakah tahun kabisat atau bukan
	if tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0) {
		fmt.Println("kabisat:", true)
	} else {
		fmt.Println("kabisat:", false)
	}
}
