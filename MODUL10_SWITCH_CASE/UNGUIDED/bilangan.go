package main

import "fmt"

func main() {
	var bilangan int

	// Input bilangan bulat dari pengguna
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)

	// Identifikasi pola dan lakukan operasi sesuai kategori
	if bilangan%10 == 0 {
		// Bilangan Kelipatan 10
		fmt.Printf("Kategori: Bilangan Kelipatan 10\nHasil pembagian antara %d / 10 = %d\n", bilangan, bilangan/10)
	} else if bilangan%5 == 0 {
		// Bilangan Kelipatan 5
		fmt.Printf("Kategori: Bilangan Kelipatan 5\nHasil kuadrat dari %d^2 = %d\n", bilangan, bilangan*bilangan)
	} else if bilangan%2 == 0 {
		// Bilangan Genap
		fmt.Printf("Kategori: Bilangan Genap\nHasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, bilangan*(bilangan+1))
	} else {
		// Bilangan Ganjil
		fmt.Printf("Kategori: Bilangan Ganjil\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, bilangan+(bilangan+1))
	}
}