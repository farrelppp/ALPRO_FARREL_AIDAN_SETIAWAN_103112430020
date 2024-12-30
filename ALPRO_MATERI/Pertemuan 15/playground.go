package main

import "fmt"

func main() {
	var membership string
	var duration int

	fmt.Scan(&membership)
	fmt.Scan(&duration)

	const tarifNormal = 65000
	const tarifKelebihan = 20000
	tarif := 0
	if duration > 2 {
		tarif = 2 * tarifNormal
		tarif += (duration - 2) * tarifKelebihan
	} else {
		tarif = duration * tarifNormal
	}
	switch membership {
	case "Gold":
		tarif = tarif / 2
	case "Silver":
		tarif = tarif * 75 / 100
	}

	fmt.Printf("IDR %d\n", tarif)
}
