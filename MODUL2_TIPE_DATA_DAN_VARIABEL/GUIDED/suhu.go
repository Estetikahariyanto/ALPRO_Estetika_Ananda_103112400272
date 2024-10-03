package main

import (
	"fmt"
)

func main() {
	var fahrenheit int

	// Meminta input dari pengguna
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Mengonversi suhu dari Fahrenheit ke Celsius
	celsius := float64(fahrenheit-32) * 5.0 / 9.0

	// Menampilkan hasil
	fmt.Printf("Suhu dalam Celsius: %.2f°C\n", celsius)
}
