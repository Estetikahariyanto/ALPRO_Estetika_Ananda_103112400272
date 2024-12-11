package main

import (
	"fmt"
)
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false 
		}
	}
	return true
}

func main() {
	var num int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}