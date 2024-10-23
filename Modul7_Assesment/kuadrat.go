package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	if N < 1 {
		fmt.Print("angka yang dimasukan tidak boleh kurang dari 1")
		return
	}

	for i := 1; i <= N; i++ {
		hasil := i * i
		fmt.Print(hasil, " ")
	}
}
