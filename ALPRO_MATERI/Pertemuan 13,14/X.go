package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("bilangan harus ganjil")
		return
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j || i+j == n+1 {
				fmt.Print(i)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
