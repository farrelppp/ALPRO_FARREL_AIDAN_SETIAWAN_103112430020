package main

import "fmt"

func main() {
	var x, y int
	var a, b bool

	fmt.Print("masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("masukkan nilai y: ")
	fmt.Scan(&y)

	xy := y % x
	yx := x % y

	a = false
	b = false
	if xy == 0 {
		a = true

	}

	if yx == 0 {
		b = true
	}

	fmt.Println(a)
	fmt.Println(b)
}
