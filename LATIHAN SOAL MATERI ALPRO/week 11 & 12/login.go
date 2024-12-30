package main

import "fmt"

func main() {
	const (
		correctUsername = "admin"
		correctPassword = "admin"
	)

	var username, password string
	var failedAttempts int

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scanln(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&password)

		if username == correctUsername && password == correctPassword {
			fmt.Println("Login berhasil")
			break
		} else {
			fmt.Println("Username atau password salah")
			failedAttempts++
		}
	}

	fmt.Printf("Jumlah gagal login: %d\n", failedAttempts)
}