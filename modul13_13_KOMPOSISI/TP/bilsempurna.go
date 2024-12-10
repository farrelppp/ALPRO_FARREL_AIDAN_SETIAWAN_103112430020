package main

import "fmt"

func angka(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func main() {
	var n int
	fmt.Scan(&n)

	if angka(n) {
		fmt.Printf("Ya %d adalah bilangan sempurna.\n", n)
	} else {
		fmt.Printf("Tidak %d bukan bilangan sempurna.\n", n)
	}
}
