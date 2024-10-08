package main

import "fmt"

func main() {
	// deklarasikan tipe data untuk variabel dibawah sebagai float64
	var jam float64
	var gaji float64
	var normal float64
	var total float64
	var lembur float64
	var bulanan float64

	normal = 40 // memberi nilai untuk variable normal dengan nilai 40

	// print perintah untuk memasukan jam kerja dan gaji
	fmt.Print("masukan jam kerja: ")
	fmt.Scan(&jam)
	fmt.Print("masukan gaji: ")
	fmt.Scan(&gaji)

	// cek apakah variabel jam lebih besar sama dengan 40
	if jam > 40 {
		// rumus perhitungan gaji bila jam lebih besar dari 40
		lembur = jam - 40
		total = (normal*gaji + lembur*1.5*gaji)
		bulanan = total * 4
		// print hasil perhitungan diatas
		fmt.Printf("Total gaji mingguan: %.2f\n", total)
		fmt.Printf("Total gaji bulanan: %.2f\n", bulanan)
	} else if jam <= 40 {
		// rumus perhitungan gaji bila jam kurang sama dengan atau sama dengan 40
		total = jam * gaji
		bulanan = total * 4
		// print hasil perhitungan diatas
		fmt.Printf("Total gaji mingguan: %.2f\n", total)
		fmt.Printf("Total gaji bulanan: %.2f\n", bulanan)
	}

}
