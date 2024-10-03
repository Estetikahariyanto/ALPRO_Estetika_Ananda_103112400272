package main

import (
	"fmt"
)

func main() {
	// Input suhu dalam Celsius dari pengguna
	var celsius float64
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scan(&celsius)

	// Mengkonversi Celsius ke Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32

	// Mengkonversi Celsius ke Reamur
	reamur := celsius * 4 / 5

	// Mengkonversi Fahrenheit ke Kelvin
	kelvin := (fahrenheit + 459.67) * 5 / 9

	// Menampilkan hasil konversi
	fmt.Printf("Celsius: %.2f\n", celsius)
	fmt.Printf("Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Reamur: %.2f\n", reamur)
	fmt.Printf("Kelvin: %.2f\n", kelvin)
}
