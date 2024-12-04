package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&input)
	if input <= 0 {
		fmt.Println("Masukan harus berupa bilangan desimal positif.")
		return
	}
	roundedUp := int(math.Ceil(input))
	sum := 0
	current := 0
	fmt.Println("Proses penjumlahan tiap iterasi:")
	for {
		current++
		sum += current
		fmt.Printf("Iterasi ke-%d: Total = %d\n", current, sum)
		if current >= roundedUp {
			break
		}
	}
	fmt.Printf("Hasil akhir: %d (dari pembulatan ke atas: %d)\n", sum, roundedUp)
}