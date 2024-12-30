package main

import "fmt"

func main() {
	var S int
	fmt.Scan(&S)

	weeks := (S + 6) / 7

	fmt.Printf("minggu ke-%d\n", weeks)
}
