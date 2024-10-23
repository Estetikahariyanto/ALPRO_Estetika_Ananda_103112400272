package main

import (
	"fmt"
)

func main() {
	var N int

	// minta input dari user
	fmt.Print("masukan nilai N: ")
	fmt.Scan(&N)

	// mencetak kuadrat dari bilangan 1 sampai N
	for i : 1; i <= N; i++ {
		fmt.Printf("Kuadrat dari %d adalah %d\n", i, i*i)
	}
 }