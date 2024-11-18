package main

import (
	"fmt"
)

func main() {
	var b int

	// Menerima input dari pengguna
	fmt.Print("Masukkan bilangan bulat (b > 1): ")
	fmt.Scan(&b)

	// Validasi input
	if b <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1.")
		return
	}

	// Mencari dan menampilkan faktor-faktor dari bilangan b
	fmt.Printf("Faktor-faktor dari %d adalah: ", b)
	faktorCount := 0
	for f := 1; f <= b; f++ {
		if b%f == 0 {
			fmt.Printf("%d ", f)
			faktorCount++ // Menghitung jumlah faktor
		}
	}
	fmt.Println()

	// Menentukan apakah bilangan tersebut adalah bilangan prima
	if faktorCount == 2 {
		fmt.Printf("%d adalah bilangan prima.\n", b)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", b)
	}
}