package main

import "fmt"

func main() {
	var celsius float64

	// Meminta input temperatur dalam derajat Celsius
	fmt.Print("Masukkan temperatur dalam Celsius: ")
	fmt.Scan(&celsius)

	// Menghitung suhu dalam Reamur
	reamur := celsius * 4 / 5

	// Menghitung suhu dalam Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32

	// Menghitung suhu dalam Kelvin
	kelvin := celsius + 273.15

	// Menampilkan hasil konversi
	fmt.Printf("Temperatur Celsius: %.2f\n", celsius)
	fmt.Printf("Derajat Reamur: %.2f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}
