package main

import "fmt"

func main() {
	var x int
	var numbers [9]int
	fmt.Scan(&x)
	for i := 0; i < 9; i++ {
		fmt.Scan(&numbers[i])
	}
	countZero := 0
	countX := 0
	for _, number := range numbers {
		if number == 0 {
			countZero++
		} else if number == x {
			countX++
		}
	}
	var modus int
	if countZero > countX {
		modus = 0
	} else {
		modus = x
	}
	fmt.Printf("Modus = %d\n", modus)
}
