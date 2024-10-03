package main

import (
	"fmt"
	"math"
)

func main() {
	var r int
	const pi = 3.1415926535

	// Meminta input jari-jari dari pengguna
	fmt.Print("Masukkan jari-jari bola (bilangan bulat): ")
	fmt.Scan(&r)

	// Menghitung volume bola
	volume := (4.0 / 3.0) * pi * math.Pow(float64(r), 3)

	// Menghitung luas permukaan bola
	luas := 4 * pi * math.Pow(float64(r), 2)

	// Menampilkan hasil perhitungan
	fmt.Printf("Volume bola: %.2f\n", volume)
	fmt.Printf("Luas permukaan bola: %.2f\n", luas)
}
