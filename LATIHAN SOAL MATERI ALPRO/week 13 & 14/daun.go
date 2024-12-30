package main
import (
	"fmt"
)

// Fungsi untuk mencari daun paling lebar
func findMaxWidth(daun []int) int {
	maxWidth := daun[0] // Inisialisasi nilai lebar maksimal
	for _, width := range daun {
		if width > maxWidth {
			maxWidth = width // Update lebar maksimal jika lebih besar
		}
	}
	return maxWidth
}
func main() {
	var n int
	fmt.Print("Masukkan jumlah daun (n): ")
	fmt.Scan(&n)

	daun := make([]int, n)
	fmt.Println("Masukkan lebar dari masing-masing daun:")
	for i := 0; i < n; i++ {
		fmt.Printf("Lebar daun ke-%d: ", i+1)
		fmt.Scan(&daun[i])
	}
	maxWidth := findMaxWidth(daun)
	fmt.Printf("Lebar daun yang paling lebar adalah: %d\n", maxWidth)
}