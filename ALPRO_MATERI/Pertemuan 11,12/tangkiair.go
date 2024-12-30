package main

import (
	"fmt"
)

func main() {
	var kapasitasTank int
	fmt.Scan(&kapasitasTank)
	totalVolume := 0
	for {
		var volumeEmber int
		fmt.Scan(&volumeEmber)
		totalVolume += volumeEmber
		if totalVolume >= kapasitasTank {
			fmt.Print(true)
			break
		} else {
			fmt.Print(false)
		}
	}
}
