package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// program menentukan angka secara acak menggunakan kodingan di bawah
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	// deklarasi variable kesempatan dan tebakan sebagai integer
	var kesempatan, tebakan int

	fmt.Print("masukan tebakan anda angka berapa yang akan muncul antara 1 sampai 100 ")
	// setiap pemain gagal menebak maka teks akan terus diprint selama kesempatan kurang atau sama dengan 5
	for kesempatan = 1; kesempatan <= 5; kesempatan++ {
		fmt.Print("masukan angka: ")
		fmt.Scan(&tebakan)

		// jika tebakan tepat sebelum kesempatan habis maka print teks dibawah
		if tebakan == target {
			fmt.Printf("angka tepat! Target: %d, Tebakan: %d, Kesempatan: %d\n", target, tebakan, kesempatan)
			break
		} else if tebakan < target {
			// jika tebakan terlalu kecil print teks dibawah
			fmt.Println("angka terlalu kecil")
		} else {
			// jika tebakan terlalu besar print teks dibawah
			fmt.Println("angka terlalu besar")
		}

	}
	// teks akan di print bila pemain gagal menebak sebanyak 5 kali atau kesempatan habis
	fmt.Printf("anda gagal menebak angka dan kehabisan kesempatan, jawabannya adalah %d.\n", target)
}
