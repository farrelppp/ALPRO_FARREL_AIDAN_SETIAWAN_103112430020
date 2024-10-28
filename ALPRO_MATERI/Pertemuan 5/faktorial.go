package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)
	// memeriksa masukan
	if n < 0 {
		fmt.Println("Bilangan harus positif!")
		return
	}

	// rumus faktorial
	faktorial := 1
	for i := 2; i <= n; i++ {
		faktorial *= i
	}
	//menampilkan hasil
	fmt.Print(faktorial)
}
