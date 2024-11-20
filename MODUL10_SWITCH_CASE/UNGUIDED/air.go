package main

import "fmt"

func main() {
	var pH float64

	// Meminta input pH dari pengguna
	fmt.Print("Masukkan kadar pH air: ")
	fmt.Scan(&pH)

	// Menentukan kelayakan air berdasarkan pH
	if pH < 0 || pH > 14 {
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	} else if pH >= 6.5 && pH <= 8.6 {
		fmt.Println("Air layak minum")
	} else {
		fmt.Println("Air tidak layak minum")
	}
}