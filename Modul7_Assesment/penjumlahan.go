package main

import "fmt"

func main() {
	var x, y, hasil int

	fmt.Scan(&x, &y)

	if x > y {
		fmt.Print("angka pertama harus lebih kecil dari angka ke dua")
		return
	}

	for i := x; i <= y; i++ {
		hasil += i
	}
	fmt.Print(hasil)

}
