package main
import (
	"fmt"
	"strconv"
)
func findLargestDigit(num int) int {
	maxDigit := 0
	numStr := strconv.Itoa(num)
	for _, digit := range numStr {
		digitValue := int(digit - '0')
		if digitValue > maxDigit {
			maxDigit = digitValue
		}
	}
	return maxDigit
}
func main() {
	var num int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&num)
	// Cek digit terbesar
	largestDigit := findLargestDigit(num)
	fmt.Printf("Digit terbesar dari bilangan %d adalah: %d\n", num, largestDigit)
}