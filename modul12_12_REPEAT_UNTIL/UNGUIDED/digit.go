package main

import "fmt"

func main() {
	var nomor int
	var continueLoop bool
	totaldigit := 0
	for continueLoop = true; continueLoop; {
		fmt.Scan(&nomor)
		if nomor > 0 {
			continueLoop = false
			digit := nomor
			for ; digit > 0; digit /= 10 {
				totaldigit++
			}
		} else {
			fmt.Print("masukan harus positif")
		}
	}
	fmt.Print(totaldigit)
}
