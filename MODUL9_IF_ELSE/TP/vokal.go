package main

import (
	"fmt"
	"strings"
)

func main() {
	var huruf string
	fmt.Print("masukan sebuah huruf: ")
	fmt.Scanln(&huruf)

	// konversi huruf ke huruf besar agar tidak case-sensitive
	huruf = strings.ToUpper(huruf)

	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("huruf vokal")
	} else {
		fmt.Println("huruf konsonan")
	}
}