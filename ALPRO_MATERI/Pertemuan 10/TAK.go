package main

import "fmt"

func main() {
	var TAK float64
	fmt.Scan(&TAK)

	if TAK < 2.00 {
		fmt.Print("Poor")
	} else if TAK >= 2.00 && TAK <= 2.75 {
		fmt.Print("Fair")
	} else if TAK >= 2.76 && TAK <= 3.00 {
		fmt.Print("Satisfactory")
	} else if TAK >= 3.01 && TAK <= 3.50 {
		fmt.Print("Very good")
	} else {
		fmt.Print("Excellent")
	}

}
