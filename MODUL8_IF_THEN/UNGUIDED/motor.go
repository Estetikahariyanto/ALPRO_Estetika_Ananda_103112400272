package main

import (
	"fmt"
)
	func main() {
		var jumlahOrang int
		fmt.Print("masukan jumlah orang: ")
		fmt.Scan(&jumlahOrang)

		// jumlah motor yang diperlukan
		jumlahMotor := (jumlahOrang + 1) / 2
		fmt.Printf("jumlah motor yang diperlukan: %\n", jumlahMotor)
	}