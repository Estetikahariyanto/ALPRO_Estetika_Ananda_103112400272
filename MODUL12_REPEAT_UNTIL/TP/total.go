package main

import (
	"fmt"
)
func main() {
	var hargaBarang, totalHarga int
	for {
		fmt.Print("Masukkan harga barang (ketik 0 jika selesai): ")
		fmt.Scanln(&hargaBarang)
		if hargaBarang == 0 {
			break
		}
		totalHarga += hargaBarang
	}
	fmt.Printf("Total harga belanja: %d\n", totalHarga)
}