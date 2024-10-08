package main

import "fmt"

func main() {
	var jamPerMinggu, upahPerJam, totalGaji float64
	const jamNormal = 40

	// Meminta input dari pengguna
	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scanln(&jamPerMinggu)

	fmt.Print("Masukkan upah per jam: ")
	fmt.Scanln(&upahPerJam)

	// Menghitung gaji mingguan
	if jamPerMinggu > jamNormal {
		lembur := jamPerMinggu - jamNormal
		totalGaji = (jamNormal * upahPerJam) + (lembur * 1.5 * upahPerJam)
	} else {
		totalGaji = jamPerMinggu * upahPerJam
	}

	// Menghitung gaji bulanan
	totalGajiBulanan := totalGaji * 4

	// Menampilkan hasil
	fmt.Printf("Total gaji bulanan karyawan adalah: %.2f\n", totalGajiBulanan)
}