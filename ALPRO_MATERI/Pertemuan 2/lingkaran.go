package main

import "fmt"

func main() {
	var (
		// Deklarasi variabel luas dan keliling lingkaran, dan pi
		luas, keliling, pi, r float32
	)
	// masukan jari jari lingkaran
	fmt.Print("masukan jari jari = ")
	fmt.Scan(&r)

	// nilai pi adalah 3.14
	pi = 3.14
	// mencari luas dan keliling lingkaran menggunakan formula luas dan keliling
	luas = pi * r * r
	keliling = 2 * pi * r
	// menampilkan hasil perhitungan luas dan keliling
	fmt.Println("luas lingkaran anda : ", luas)
	fmt.Println("keliling lingkaran anda : ", keliling)

}
