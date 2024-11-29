package main

import (
	"fmt"
)

func main() {
	var x, y int

	// Meminta input dua bilangan bulat positif
	fmt.Print("Masukkan bilangan x (x >= y): ")
	_, errX := fmt.Scan(&x)
	fmt.Print("Masukkan bilangan y (y