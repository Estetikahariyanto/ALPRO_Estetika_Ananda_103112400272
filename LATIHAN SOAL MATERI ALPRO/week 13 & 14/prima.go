package main
import (
	"fmt"
)
// Fungsi ini untuk ngecek bilangan prima atau bukan
func isPrime(n int) (bool, []int) {
	if n < 2 {
		return false, nil
	}
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i) // Menyimpan faktor dari bilangan
		}
	}
	return len(factors) == 2, factors
}
func main() {
	var x int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&x)
	// Cek apakah bilangan prima
	isPrimeNumber, factors := isPrime(x)
	if isPrimeNumber {
		fmt.Printf("%d adalah bilangan prima\n", x)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", x)
		fmt.Printf("Faktor dari %d adalah: %v\n", x, factors)
	}
}