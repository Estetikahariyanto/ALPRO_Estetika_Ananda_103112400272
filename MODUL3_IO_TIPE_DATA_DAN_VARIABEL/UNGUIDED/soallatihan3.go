package main

import "fmt"

func isLeapYear(year int) bool {
	// Tahun kabisat habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var year int

	// Meminta input tahun dari pengguna
	fmt.Print("2021: "
