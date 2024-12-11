package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung f(K)
func calculateF(k int) float64 {
	numerator := math.Pow(float64(4*k+2), 2)
	denominator := float64((4*k+1)*(4*k+3))
	return numerator / denominator
}

// Fungsi untuk menghitung akar 2 dengan hampiran hingga K
func calculateSqrt2(k int) float64 {
	result := 1.0
	for i := 0; i <= k; i++ {
		result *= calculateF(i)
	}
	return result
}

func main() {
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	// Menghitung nilai f(K)
	fk := calculateF(k)
	fmt.Printf("Nilai f(K) = %.10f\n", fk)

	// Menghitung hampiran akar 2
	sqrt2 := calculateSqrt2(k)
	fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2)
}