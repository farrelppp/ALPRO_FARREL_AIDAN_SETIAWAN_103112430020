package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	a += (a % 2) ^ 1

	n := (b-a)/2 + 1
	sum := n * (2*a + (n-1)*2) / 2

	fmt.Print(sum)
}
