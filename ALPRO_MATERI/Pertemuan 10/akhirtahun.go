package main

import (
	"fmt"
)

func main() {
	
	var totalBelanja int
	var bersediaKartu bool

	fmt.Scan(&totalBelanja)
	fmt.Scan(&bersediaKartu)

	
	var memperolehKartu, memperolehDiskon, memperolehCashback bool
	finalTotal := totalBelanja

	
	if bersediaKartu {
		memperolehKartu = true
	}

	if totalBelanja >= 100000 {
		memperolehDiskon = true
		diskon := totalBelanja / 10
		finalTotal -= diskon
		if totalBelanja >= 200000 {
			if memperolehKartu {
				memperolehCashback = true
				finalTotal -= 75000
			}
		}
	} else if totalBelanja < 100000 {
	
		memperolehDiskon = false
		memperolehCashback = false
	}

	
	fmt.Printf("Kartu? %v\n", memperolehKartu)
	fmt.Printf("Diskon? %v\n", memperolehDiskon)
	fmt.Printf("Cashback? %v\n", memperolehCashback)
	fmt.Printf("Rp. %d\n", finalTotal)
}
