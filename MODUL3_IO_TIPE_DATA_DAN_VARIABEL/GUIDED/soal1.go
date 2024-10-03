package main

import (
	"fmt"
)

func main() {
	// Panjang sisi alun-alun
	var sisi float64 = 27

	// Menghitung keliling
	keliling := 4 * sisi

	// Menghitung luas
	luas := sisi * sisi

	// Menampilkan hasil
	fmt.Printf("Keliling alun-alun: %.2f meter\n", keliling)
	fmt.Printf("Luas alun-alun: %.2f meter persegi\n", luas)
}
