package main

import (
	"fmt"
)
func main() {
	var N int
	fmt.Print("N: ")
	fmt.Scan(&N)
	if N <= 0 {
		fmt.Println("Pita: ")
		return
	}
	pita := ""
	for i := 0; i < N; i++ {
		var bunga string
		fmt.Printf("Bunga %d: ", i+1)
		fmt.Scan(&bunga)

		if pita == "" {
			pita = bunga
		} else {
			pita += " - " + bunga
		}
	}
	fmt.Println("Pita:", pita)
}