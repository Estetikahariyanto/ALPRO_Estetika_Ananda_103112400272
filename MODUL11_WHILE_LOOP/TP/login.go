package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const password = "estetika"
	const maxAttempts = 3

	reader := bufio.NewReader(os.Stdin)
	attempts := 0

	for attempts < maxAttempts {
		fmt.Print("Masukkan password: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi kesalahan saat membaca input.")
			continue
		}

		// Menghapus karakter newline dari input
		input = strings.TrimSpace(input)

		if input == password {
			fmt.Println("Login berhasil!")
			return
		} else {
			fmt.Println("Password salah.")
			attempts++
		}
	}

	fmt.Println("Login ditolak.")
}