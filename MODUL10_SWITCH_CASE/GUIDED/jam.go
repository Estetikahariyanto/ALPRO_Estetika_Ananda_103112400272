package main

import "fmt"

func main() {
	var jam24 int

	// Input jam dalam format 24 jam
	fmt.Print("Masukkan jam (0-23): ")
	fmt.Scan(&jam24)

	// Konversi ke format 12 jam
	jam12 := jam24 % 12
	if jam12 == 0 {
		jam12 = 12
	}

	periode := "AM"
	if jam24 >= 12 {
		periode = "PM"
	}

	// Menampilkan hasil
	fmt.Printf("%d dalam format 24 jam adalah %d %s\n", jam24, jam12, periode)
}