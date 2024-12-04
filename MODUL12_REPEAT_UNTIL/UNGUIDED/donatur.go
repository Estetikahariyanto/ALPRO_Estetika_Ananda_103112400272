package main

import "fmt"

func main() {
	var target, totalDonasi, donasi int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)
	fmt.Println("Masukkan donasi (masukan dihentikan jika target tercapai):")
	for {
		fmt.Print("Donasi: ")
		fmt.Scan(&donasi)
		totalDonasi += donasi
		if totalDonasi >= target {
			break
		}
	}
	fmt.Printf("Target tercapai! Total donasi: %d\n", totalDonasi)
}