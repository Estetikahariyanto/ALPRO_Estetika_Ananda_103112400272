package main

import (
	"fmt"
	"strings"
)

func main() {
	var umur int
	var kewarganegaraan string

	// input umur nya
	fmt.Print("Masukkan umur Anda: ")
	fmt.Scanln(&umur)

	// input kewarganegaraan
	fmt.Print("Masukkan kewarganegaraan Anda (contoh: WNI atau WNA): ")
	fmt.Scanln(&kewarganegaraan)

	// konversi kewarganegaraan ke huruf besar
	kewarganegaraan = strings.ToUpper(kewarganegaraan)

	// cek syarat untuk mengikuti Pemilu
	if umur >= 17 && kewarganegaraan == "WNI" {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else {
		// menampilkan alasan tidak bisa mengikuti pemilu
		if umur < 17 && kewarganegaraan != "WNI" {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena belum berusia 17 tahun dan bukan WNI.")
		} else if umur < 17 {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena belum berusia 17 tahun.")
		} else if kewarganegaraan != "WNI" {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena bukan WNI.")
		}
	}
}
