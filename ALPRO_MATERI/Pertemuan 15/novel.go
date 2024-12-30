package main

import "fmt"

func main() {
	var P, R int
	fmt.Scan(&P)
	fmt.Scan(&R)
	days := (P + R - 1) / R
	fmt.Printf("%d hari\n", days)
}
