package main

import (
	"fmt"
)

func main() {
	// Meminta input usia dari pengguna
	var usia int
	fmt.Print("Masukkan usia Anda: ")
	fmt.Scanln(&usia)

	// Menentukan kategori usia menggunakan switch
	var kategori string
	switch {
	case usia >= 0 && usia <= 12:
		kategori = "Anak-anak"
	case usia >= 13 && usia <= 17:
		kategori = "Remaja"
	case usia >= 18 && usia <= 64:
		kategori = "Dewasa"
	case usia >= 65:
		kategori = "Lansia"
	default:
		fmt.Println("Usia tidak valid")
		return
	}

	// Menampilkan hasil kategori usia
	fmt.Printf("Kategori usia Anda: %s\n", kategori)
}