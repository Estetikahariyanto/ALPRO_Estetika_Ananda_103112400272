package main

import (
	"fmt"
)

func main() {
	// Meminta input dari pengguna
	var n int
	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&n)

	// Mengecek apakah input valid
	if n <= 0 {
		fmt.Println("Harap masukkan bilangan bulat positif.")
		return
	}

	// Menghitung jumlah deret angka dari 1 hingga n
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	// Menampilkan hasil penjumlahan
	fmt.Printf("Jumlah deret angka dari 1 hingga %d adalah: %d\n", n, sum)
}
