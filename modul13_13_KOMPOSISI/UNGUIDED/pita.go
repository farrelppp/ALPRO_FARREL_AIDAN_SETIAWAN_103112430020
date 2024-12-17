package main

import (
	"fmt"
	"strings"
)

func main() {
	var bunga string
	var input string
	count := 0
	for {
		count++
		fmt.Printf("Bunga %d : ", count)
		fmt.Scan(&input)

		if input == "SELESAI" {

			if bunga == "" {
				fmt.Print("Pita :")
				fmt.Print("Bunga: 0")
			} else {
				fmt.Print("Pita:", bunga)
				fmt.Print("Bunga:", len(strings.Split(bunga, " - ")))
			}
			break
		}
		if bunga != "" {
			bunga += input + " - "
		} else {
			bunga = input
		}
	}
}
