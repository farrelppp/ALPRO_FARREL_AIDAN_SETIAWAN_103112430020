package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64

	for continueLoop := true; continueLoop; {
		fmt.Scan(&input)
		bulatkeatas := math.Ceil(input)
		for i := input; i <= bulatkeatas; i = math.Round((i+0.1)*10) / 10 {
			fmt.Printf("%.1f\n", i)
		}
		continueLoop = false
	}
}
