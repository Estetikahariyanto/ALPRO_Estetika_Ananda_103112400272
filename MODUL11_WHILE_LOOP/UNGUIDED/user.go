package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const correctUsername = "Admin"
	const correctPassword = "Admin"
	var failedAttempts int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Program Login ===")

	for {
		fmt.Print("Masukkan username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Masukkan password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		if username == correctUsername && password == correctPassword {
			fmt.Printf("Login berhasil! %d percobaan gagal login.\n", failedAttempts)
			break
		} else {
			fmt.Println("Username atau password salah. Silakan coba lagi.")
			failedAttempts++
		}
	}
}