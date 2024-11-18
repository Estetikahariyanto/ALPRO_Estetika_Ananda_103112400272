package main

import (
	"fmt"
)

func main() {

	var nilai int
	fmt.Print("masukan nilai mahasiswa")
	fmt.Scanln(&nilai)

	indeks := klasifikasiNilai(nilai)
	fmt.Printf("nilai %d mendapatkan indeks %s\n", nilai, indeks)
}

func klasifikasiNilai(nilai int) string {
	if nilai > 90 {
		return "A"
	} else if nilai >= 80 && nilai <=90 {
		return "AB"
	} else if nilai >= 70 && nilai < 80 {
		return "B"
	} else {
		return "C"
	}
}