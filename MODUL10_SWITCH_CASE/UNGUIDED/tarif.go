package main

import (
	"fmt"
	"math"
)

func main() {
	var jenisKendaraan string
	var durasiParkir float64

	// Input jenis kendaraan dan durasi parkir
	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jenisKendaraan)

	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&durasiParkir)

	// Membulatkan durasi ke atas jika kurang dari 1 jam
	jamParkir := math.Ceil(durasiParkir)

	// Menentukan tarif parkir
	tarifPerJam := 0
	if jenisKendaraan == "motor" {
		tarifPerJam = 2000
	} else if jenisKendaraan == "mobil" {
		tarifPerJam = 5000
	} else if jenisKendaraan == "truk" {
		tarifPerJam = 8000
	} else {
		fmt.Println("Jenis kendaraan tidak valid.")
		return
	}

	// Menghitung total biaya parkir
	totalBiaya := tarifPerJam * int(jamParkir)
	fmt.Printf("Total biaya parkir: Rp %d\n", totalBiaya)
}