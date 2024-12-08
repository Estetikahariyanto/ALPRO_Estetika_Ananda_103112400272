package main

import (
	"fmt"
)
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println("Masukkan bilangan bulat yang lebih besar dari 0.")
		return
	}
	fmt.Printf("Bilangan prima dari 1 hingga %d:\n", n)
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}