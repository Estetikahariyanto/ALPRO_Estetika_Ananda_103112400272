package main

import (
	"fmt"
)

func hitungBiayaPengiriman(beratGram int) int {
	// Konversi berat ke kg dan sisa gram
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	// Hitung biaya dasar (Rp 10.000 per kg)
	biayaDasar := beratKg * 10000

	// Hitung biaya tambahan untuk sisa gram
	var biayaTambahan int
	if beratKg > 10 {
		biayaTambahan = 0 // Gratis untuk sisa berat jika berat lebih dari 10 kg
	} else {
		if sisaGram >= 500 {
			biayaTambahan = sisaGram * 5 // Rp 5 per gram jika sisa >= 500 gram
		} else {
			biayaTambahan = sisaGram * 15 // Rp 15 per gram jika sisa < 500 gram
		}
	}

	// Total biaya
	totalBiaya := biayaDasar + biayaTambahan
	return totalBiaya
}

func main() {
	// Input berat parsel dalam gram
	var beratParsel int
	fmt.Print("Masukkan berat parsel dalam gram: ")
	fmt.Scan(&beratParsel)

	// Hitung total biaya
	totalBiaya := hitungBiayaPengiriman(beratParsel)

	// Tampilkan total biaya
	fmt.Printf("Total biaya pengiriman: Rp %d\n", totalBiaya)
}