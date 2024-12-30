package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// Membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan satu karakter: ")
	input, _ := reader.ReadString('\n')

	// Menghapus karakter newline
	char := rune(input[0])

	// Memeriksa apakah karakter adalah alfabet
	if unicode.IsLetter(char) {
		fmt.Println("Alphabet")
	} else {
		fmt.Println("Bukan Alphabet")
	}
}
