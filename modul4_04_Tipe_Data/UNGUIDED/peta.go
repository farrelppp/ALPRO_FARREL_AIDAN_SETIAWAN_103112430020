package main

import (
	"fmt"
	"math"
)

func hitung(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

}

func main() {
	var ax, ay, bx, by, cx, cy float64

	fmt.Println("masukan koordinat titik A: ")
	fmt.Scan(&ax, &ay)
	fmt.Println("masukan koordinat titik B: ")
	fmt.Scan(&bx, &by)
	fmt.Println("masukan koordinat titik C: ")
	fmt.Scan(&cx, &cy)

	AB := hitung(ax, ay, bx, by)
	BC := hitung(bx, by, cx, cy)
	CA := hitung(cx, cy, ax, ay)

	terpanjang := AB
	if BC > terpanjang {
		terpanjang = BC
	}
	if CA > terpanjang {
		terpanjang = CA
	}

	// Menampilkan hasil
	fmt.Printf("sisi terpanjang adalah: %.2f\n", terpanjang)

}
