package main

import "fmt"

func main() {
	var t [5]float64
	fmt.Println("Masukkan 5 catatan temperatur (pisahkan dengan spasi):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&t[i])
	}

	stableUp := true
	stableDown := true

	for i := 1; i < 5; i++ {
		if t[i] > t[i-1] {
			stableDown = false
		} else if t[i] < t[i-1] {
			stableUp = false
		} else {
			stableUp = false
			stableDown = false
		}
	}

	if stableUp {
		fmt.Println("Stabil naik")
	} else if stableDown {
		fmt.Println("Stabil turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}