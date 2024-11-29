package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Masukkan bilangan bulat positif: ")
	_, err := fmt.Scan(&number)

	// Validasi masukan harus bilangan bulat positif
	if err != nil || number <= 0 {
		fmt.Println("Masukan tidak valid. Harap masukkan bilangan bulat positif.")
		return
	}

	fmt.Println("Digit-digit bilangan tersebut dari kanan ke kiri:")

	// Loop untuk mencacah digit dari kanan ke kiri
	for number > 0 {
		digit := number % 10 // Mendapatkan digit terakhir
		fmt.Println(digit)   // Mencetak digit terakhir
		number /= 10         // Menghilangkan digit terakhir
	}
}