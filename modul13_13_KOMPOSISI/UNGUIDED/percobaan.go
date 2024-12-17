package main

import (
	"fmt"
)

func main() {
	masukan := []string{"merah", "kuning", "hijau", "ungu"} //merah kuning hijau ungu
	benar := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		var a, b, c, d string
		fmt.Scan(&a, &b, &c, &d)
		if a != masukan[0] || b != masukan[1] || c != masukan[2] || d != masukan[3] {
			benar = false
		}
	}

	fmt.Printf("BERHASIL: %t\n", benar)
}
