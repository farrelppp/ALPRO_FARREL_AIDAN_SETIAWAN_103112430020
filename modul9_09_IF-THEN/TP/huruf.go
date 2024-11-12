package main

import "fmt"

func main() {
	var huruf string
    // perintah untuk memasukan huruf
	fmt.Print("masukan huruf: ")
	fmt.Scan(&huruf)

	// fungsi if untuk menentukan apakah huruf vokal atau konsonan
	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Print("masukan adalah huruf vokal")
	} else {
		fmt.Print("masukan adalah huruf konsonan")
	}
}
