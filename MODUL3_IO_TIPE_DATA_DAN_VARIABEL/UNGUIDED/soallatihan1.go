package main

import (
	"fmt"
)

func calculateX(fx float64) (float64, error) {
	// f(x) = (2 / (x + 5)) + 5
	// Rearranging the equation:
	// fx - 5 = 2 / (x + 5)
	// (fx - 5) * (x + 5) = 2
	// x + 5 = 2 / (fx - 5)
	// x = (2 / (fx - 5)) - 5

	if fx <= 5 {
		return 0, fmt.Errorf("nilai f(x) harus lebih besar dari 5 untuk menghindari pembagian dengan nol")
	}

	x := (2 / (fx - 5)) - 5
	return x, nil
}

func main() {
	var fx float64

	// Memasukkan nilai f(x)
	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)

	// Menghitung nilai x
	x, err := calculateX(fx)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Nilai x adalah: %.2f\n", x)
	}
}
