package main

import "fmt"

func main() {
	var token string
	maxPercobaan := 3
	percobaan := 0

	for percobaan < maxPercobaan {
		fmt.Print("Masukkan password: ")
		fmt.Scan(&token)

		if token == "12345abcde" {
			fmt.Println("Selamat, Anda berhasil login!")
			return
		}

		percobaan++
		if percobaan < maxPercobaan {
			fmt.Println("Maaf, password salah")
		}
	}
	fmt.Println("Anda tidak bisa login, kesempatan anda sudah 3x login habis!")
}
