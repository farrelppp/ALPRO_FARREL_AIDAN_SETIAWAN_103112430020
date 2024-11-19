package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	// perintah untuk memasukan nilai akhir
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	// Tentukan nmk berdasarkan nam
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	// Output hasil penilaian
	fmt.Println("Nilai mata kuliah:", nmk)
}

// a. keluaran tidak ada karena variable nmk tidak diubah
//    eksekusi tidak sesuai spesifikasi karena ada kesalahan program
// b. - kesalahan tipe data
//      variable nam seharusnya angka namun malah string
//      variable nmk tidak terisi nilai apapun sehingga output tidak ada
//    - kesalahan if
//      beberapa program ifnya terpisah
//    - alur yang benar
//      menggunakan if-else if
//      nam float untuk bilangan/angka dan nmk string untuk huruf
