package main

import (
	"fmt"
)

func calculateTotalAfterDiscount(initialTotal, discount int) float64 {
	discountAmount := float64(initialTotal) * (float64(discount) / 100.0)
	finalTotal := float64(initialTotal) - discountAmount
	return finalTotal
}

func main() {
	var initialTotal, discount int

	// Input total inisial pengeluaran dan jumlah diskon
	fmt.Print("Masukkan total inisial pengeluaran: ")
	fmt.Scan(&initialTotal)
	fmt.Print("Masukkan jumlah diskon (%): ")
	fmt.Scan(&discount)

	// Menghitung total setelah diskon
	finalTotal := calculateTotalAfterDiscount(initialTotal, discount)

	// Menampilkan hasil
	fmt.Printf("Total setelah diskon: %.2f\n", finalTotal)
}
