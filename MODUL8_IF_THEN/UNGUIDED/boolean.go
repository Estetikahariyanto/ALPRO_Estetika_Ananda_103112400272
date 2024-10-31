package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("masukan bilangan x: ")
	fmt.Scan(&x)
	fmt.Scan("masukan bilangan y: ")
	fmt.Scan(&y)

	// ngecek x itu faktor dari y atau bukan
	faktorXY := y%x == 0

	// ngecek y itu faktor dari x atau bukan
	faktorYX := x%y == 0

	// menampilkan hasil
	fmt.Println(faktorXY)
	fmt.Println(faktorYX)
}