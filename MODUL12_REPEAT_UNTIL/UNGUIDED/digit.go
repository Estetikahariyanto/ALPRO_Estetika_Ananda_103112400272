package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)
	if number <= 0 {
		fmt.Println("Masukan harus berupa bilangan bulat positif.")
		return
	}
	digitCount := 0
	temp := number
	for {
		temp /= 10
		digitCount++
		if temp == 0 {
			break
		}
	}
	fmt.Printf("Banyaknya digit dari bilangan %d adalah: %d\n", number, digitCount)
}