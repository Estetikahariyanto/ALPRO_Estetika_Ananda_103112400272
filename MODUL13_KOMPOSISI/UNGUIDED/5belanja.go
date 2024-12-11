package main

import (
	"fmt"
)
func main() {
	for {
		var beratKiri, beratKanan float64
		fmt.Print("Masukkan berat belanjaan di kedua kantong (misalnya: 5.5 10): ")
		fmt.Scan(&beratKiri, &beratKanan)
		if beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Berat tidak boleh negatif. Program selesai.")
			break
		}
		if beratKiri >= 9 || beratKanan >= 9 {
			fmt.Println("Proses selesai.")
			break
		}
		totalBerat := beratKiri + beratKanan
		if totalBerat > 150 {
			fmt.Println("Total berat melebihi 150 kg. Program selesai.")
			break
		}
		selisihBerat := beratKiri - beratKanan
		if selisihBerat < 0 {
			selisihBerat = -selisihBerat
		}
		if selisihBerat >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}