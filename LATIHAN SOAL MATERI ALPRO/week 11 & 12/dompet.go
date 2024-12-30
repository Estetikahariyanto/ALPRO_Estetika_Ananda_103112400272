package main

import "fmt"

func main() {
	var saldo, transaksi int

	fmt.Println("Masukkan transaksi (masukkan 0 untuk selesai):")
	for {
		fmt.Scan(&transaksi)
		if transaksi == 0 {
			break
		}
		saldo += transaksi
	}

	fmt.Printf("Jumlah saldo akhir: %d\n", saldo)
}