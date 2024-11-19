package main

import "fmt"

func main() {
	var input int

	fmt.Scan(&input)

	if input < 1000 || input > 9999 {
		fmt.Print("inputan adalah angka 4 digit")
		return

	}

	a1 := input / 1000
	b2 := input % 1000 / 100
	c3 := input % 100 / 10
	d4 := input % 10

	if a1 < b2 && b2 < c3 && c3 < d4 {
		fmt.Print("digit pada ", input, " terurut membesar")
	} else if a1 > b2 && b2 > c3 && c3 > d4 {
		fmt.Print("digit pada ", input, " terurut mengecil")
	} else {
		fmt.Print("digit pada ", input, " tidak terurut")
	}

}
