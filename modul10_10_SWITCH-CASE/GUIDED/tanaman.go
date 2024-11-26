package main

import "fmt"

func main() {
	var tanaman string
	fmt.Scan(&tanaman)

	switch tanaman{
	case "nepenthes", "drosera":
		fmt.Print("termasuk tanaman karniovora")
		fmt.Print("asli indonesia")
	
	case "venus", "saracenia":
		fmt.Print("termasuk tanaman karniovora")
		fmt.Print("tidak asli indonesia")
	default :
		fmt.Print("bukan tanaman karniovora")
	}
}