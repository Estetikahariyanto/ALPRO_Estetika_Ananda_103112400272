package main

import "fmt"

func main() {
	var jenisKendaraan string
	var durasi float64
	var tarif int

	// Input
	fmt.Print("Jenis kendaraan (motor/mobil/truk): ")
	fmt.Scanln(&jenisKendaraan)
	fmt.Print("Durasi parkir (jam): ")
	fmt.Scanln(&durasi)

	// Penentuan tarif
	switch jenisKendaraan {
	case "motor":
		tarif = 2000
	case "mobil":
		tarif = 5000
	case "truk":
		tarif = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid.")
		return
	}

	// Pembulatan durasi kurang dari 1 jam
	if durasi < 1 {
		durasi = 1
	}

	// Perhitungan total biaya
	total := int(durasi) * tarif

	// Output
	fmt.Printf("Total biaya parkir: Rp %d\n", total)
}