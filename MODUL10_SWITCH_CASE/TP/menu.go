package main

import (
	"fmt"
)

func main() {
	// Menampilkan daftar menu restoran
	fmt.Println("Menu Restoran Cepat Saji")
	fmt.Println("1. Burger - Rp25,000")
	fmt.Println("2. Fried Chicken - Rp20,000")
	fmt.Println("3. French Fries - Rp15,000")
	fmt.Println("4. Soft Drink - Rp10,000")
	fmt.Println("5. Coffee - Rp15,000")

	// Meminta pengguna memasukkan kode menu
	var kode int
	fmt.Print("Masukkan kode menu (1-5): ")
	fmt.Scanln(&kode)

	// Menggunakan switch untuk menentukan pilihan menu
	var namaItem string
	var harga int
	switch kode {
	case 1:
		namaItem = "Burger"
		harga = 25000
	case 2:
		namaItem = "Fried Chicken"
		harga = 20000
	case 3:
		namaItem = "French Fries"
		harga = 15000
	case 4:
		namaItem = "Soft Drink"
		harga = 10000
	case 5:
		namaItem = "Coffee"
		harga = 15000
	default:
		fmt.Println("Kode menu tidak valid")
		return
	}

	// Menampilkan pilihan menu
	fmt.Printf("Anda memilih %s - Rp%d\n", namaItem, harga)
}