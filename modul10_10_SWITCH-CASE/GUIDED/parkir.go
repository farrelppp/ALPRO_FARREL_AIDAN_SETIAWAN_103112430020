package main

import "fmt"

func main() {
	var jenis string
	var tarif, durasi int

	fmt.Scan(&jenis)
	fmt.Scan(&durasi)

	switch {
	case jenis == "motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case jenis == "motor" && durasi > 2:
		tarif = 9000
	case jenis == "mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case jenis == "mobil" && durasi > 2:
		tarif = 20000
	case jenis == "truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000
	case jenis == "truk" && durasi > 2:
		tarif = 35000
	default:
		fmt.Print("Jenis kendaraan atau durasi waktu parkir tidak valid")
	}
	fmt.Print(tarif)
}
