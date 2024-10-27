package main

import (
	"fmt"
)

func main() {
	var N int

	// Pengguna minta untuk memasukkan bilangan bulat positif
	fmt.Print("Masukkan bilangan bulat positif N: ")
	_, err := fmt.Scan(&N)
	if err != nil || N <= 0 {
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat positif.")
		return
	}

	// Menampilkan N
	for d := 1; d <= N; d++ {
		isFactor := N%d == 0 // Mengecek apakah d adalah faktor dari N
		fmt.Printf("%d %t\n", d, isFactor) // Menampilkan d dan apakah d adalah faktor
	}