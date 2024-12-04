package main

import (
	"fmt"
	"strings"
)
func main() {
	var userInput string
	for {
		fmt.Print("Masukkan kata (ketik 'telkom' untuk berhenti): ")
		fmt.Scanln(&userInput)
		if strings.ToLower(userInput) == "telkom" {
			fmt.Println("Program selesai.")
			break
		}
	}
}