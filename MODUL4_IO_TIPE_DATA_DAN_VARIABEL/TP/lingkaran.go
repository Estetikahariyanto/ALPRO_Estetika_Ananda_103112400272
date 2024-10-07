package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	const pi = math.Pi

	// Meminta pengguna memasukkan jari-jari
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)

	// Menghitung luas dan keliling lingkaran
	luas := pi * r * r
	keliling := 2 * pi * r

	// Menampilkan hasil
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah: %.2f\n", r, luas)
	fmt.Printf("Keliling lingkaran dengan jari-jari %.2f adalah: %.2f\n", r, keliling)
}
