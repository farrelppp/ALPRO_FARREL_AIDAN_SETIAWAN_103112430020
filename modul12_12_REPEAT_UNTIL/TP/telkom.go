package main

import "fmt"

func main() {
	var masukan string

	for {
		fmt.Print("Masukan kata: ")
		fmt.Scan(&masukan)

		if masukan == "telkom" || masukan == "Telkom" || masukan == "TELKOM" {
			fmt.Println("program selesai")
			break
		} else {
			fmt.Println("Anda mengetik:", masukan)

		}

	}
}
