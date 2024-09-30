package main

import "fmt"

func main() {
	var nama, nim, kelas string
	fmt.Println("masukan nama:")
	fmt.Scanln(&nama)
	fmt.Println("masukan nim:")
	fmt.Scanln(&nim)
	fmt.Println("masukan kelas:")
	fmt.Scanln(&kelas)
	fmt.Println("Perkenalkan saya adalah", nama, ", salah satu mahasiswa Prodi S1-IF dari kelas", kelas, ",dengan NIM", nim)
}
