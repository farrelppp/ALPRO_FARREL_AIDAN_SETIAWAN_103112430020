package main

import "fmt"

func main() {
	var huruf string

	fmt.Print("masukan huruf: ")
	fmt.Scan(&huruf)

	if huruf >= "a" && huruf <= "z" {
		fmt.Println("harus dalam format capslock")

	} else if !(huruf >= "A" && huruf <= "Z") && !(huruf >= "a" && huruf <= "z") {
		fmt.Println("Masukan huruf alfabet")
		return
	}

	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("Huruf ini adalah huruf vokal")
	} else if huruf != "A" || huruf != "I" || huruf != "U" || huruf != "E" || huruf != "O" {
		fmt.Println("Huruf ini adalah huruf konsonan")
	}
}
