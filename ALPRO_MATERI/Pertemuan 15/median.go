package main

import (
	"fmt"
	"sort"
)

func main() {
	var y int
	var numbers [9]int

	fmt.Scan(&y)

	for i := 0; i < 9; i++ {
		fmt.Scan(&numbers[i])
	}

	sort.Ints(numbers[:])

	median := numbers[4]

	fmt.Printf("Median adalah: %d\n", median)
}
