package main

import "fmt"

func main() {
	// loop untuk bilangan genap 1-50
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
