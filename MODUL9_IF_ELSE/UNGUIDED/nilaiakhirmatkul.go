package main

import (
	"fmt"
)

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	// Tentukan NMK berdasarkan rentang nilai yang sesuai
	if nam >= 88 && nam <= 100 {
		nmk = "A"
	} else if nam >= 72.5 && nam < 88 {
		nmk = "AB"
	} else if nam >= 65 && nam < 72.5 {
		nmk = "B"
	} else if nam >= 57.5 && nam < 65 {
		nmk = "BC"
	} else if nam >= 50 && nam < 57.5 {
		nmk = "C"
	} else if nam >= 40 && nam < 50 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah:", nmk)
}
