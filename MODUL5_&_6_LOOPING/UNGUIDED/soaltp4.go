package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inisialisasi seed untuk random number generator
	rand.Seed(time.Now().UnixNano())

	// Memilih angka acak antara 1 hingga 100
	secretNumber := rand.Intn(100) + 1
	var guess int
	maxAttempts := 5

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Saya telah memilih sebuah angka antara 1 hingga 100.")
	fmt.Printf("Anda memiliki %d kesempatan untuk menebak.\n", maxAttempts)

	// Memberikan 5 kesempatan kepada pengguna untuk menebak angka
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		fmt.Printf("Percobaan %d: Masukkan tebakan Anda: ", attempts)
		fmt.Scan(&guess)

		// Mengecek apakah tebakan benar
		if guess == secretNumber {
			fmt.Println("Selamat! Anda berhasil menebak angka dengan benar.")
			return
		} else if guess < secretNumber {
			fmt.Println("Terlalu kecil!")
		} else {
			fmt.Println("Terlalu besar!")
		}
	}

	// Jika pengguna tidak berhasil menebak dalam 5 kali percobaan
	fmt.Printf("Maaf, Anda tidak berhasil menebak angkanya. Angka yang benar adalah %d.\n", secretNumber)
}
