package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	if N < 1 {
		fmt.Print("angka yang dimasukan harus bulat positif")
		return
	}

	for i := 1; i <= N; i++ {
		hasil := i * i
		fmt.Print(hasil, " ")
	}
}
