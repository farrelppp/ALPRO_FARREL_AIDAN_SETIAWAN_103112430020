package main

import "fmt"

func main() {
	var dinar, dirham, fals, qirat int

	fmt.Scan(&qirat)

	dinar = qirat / 600 % 10
	dirham = qirat / 60 % 10
	fals = qirat / 6 % 10
	qirat = qirat % 6

	fmt.Printf("%d %d %d %d \n", dinar, dirham, fals, qirat)

}
