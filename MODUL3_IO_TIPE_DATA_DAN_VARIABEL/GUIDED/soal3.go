package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64

	// Meminta input suhu dalam Fahrenheit dari pengguna
	fmt.Print("Masukkan suhu dalam derajat Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Mengonversi Fahrenheit ke Kelvin
	kelvin := (fahrenheit-32)*5/9 + 273.15

	// Menampilkan hasil konversi
	fmt.Printf("Suhu dalam derajat Kelvin: %.2f K\n", kelvin)
}
