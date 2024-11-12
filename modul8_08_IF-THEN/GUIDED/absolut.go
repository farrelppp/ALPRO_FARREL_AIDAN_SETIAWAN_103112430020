package main

import "fmt"

func main() {
	var bilangan int

	fmt.Scan(&bilangan)

	if bilangan < 0 {
		bilangan = -bilangan //-(-3) = +3
	}
	fmt.Print(bilangan)
}