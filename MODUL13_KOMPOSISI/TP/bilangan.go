package main

import (
	"fmt"
)
func sumOfFactors(num int) int {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum
}
func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Masukkan bilangan positif.")
		return
	}
	if sumOfFactors(n) == n {
		fmt.Printf("Ya, %d adalah bilangan sempurna.\n", n)
	} else {
		fmt.Printf("Tidak, %d bukan bilangan sempurna.\n", n)
	}
}