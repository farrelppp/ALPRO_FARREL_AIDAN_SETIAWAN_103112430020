package main

import "fmt"

func main() {
	var umur int
	var wargaNegara string
	// perintah untuk memasukan umur
	fmt.Print("masukkan umur: ")
	fmt.Scan(&umur)
	
	// perintah untuk memasukan kewarganegaraan sesuai ketentuan
	fmt.Print("masukkan kewarganegaraan (CAPSLOCK): ")
	fmt.Scan(&wargaNegara)

	// fungsi untuk menentukan apakah warga bisa mengikuti pemilu atau tidak
	if umur >= 17 && wargaNegara == "INDONESIA" {
		fmt.Print("Anda bisa mengikuti pemilu")

	} else if umur < 17 || wargaNegara != "INDONESIA" {
		fmt.Print("Anda belum bisa mengikuti pemilu")
	}
}
