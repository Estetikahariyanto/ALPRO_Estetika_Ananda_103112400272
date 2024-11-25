package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var totalBelanja float64 = 0
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Selamat datang di aplikasi kasir sederhana!")
	for {
		fmt.Println("\nPilih aksi:")
		fmt.Println("1. Tambah barang")
		fmt.Println("2. Selesaikan transaksi")

		fmt.Print("Masukkan pilihan (1/2): ")
		input, _ := scanner.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			// Tambah barang
			fmt.Print("Masukkan nama barang: ")
			namaBarang, _ := scanner.ReadString('\n')
			namaBarang = strings.TrimSpace(namaBarang)

			fmt.Print("Masukkan harga barang: ")
			hargaBarangStr, _ := scanner.ReadString('\n')
			hargaBarangStr = strings.TrimSpace(hargaBarangStr)

			hargaBarang, err := strconv.ParseFloat(hargaBarangStr, 64)
			if err != nil {
				fmt.Println("Input harga tidak valid, coba lagi.")
				continue
			}

			totalBelanja += hargaBarang
			fmt.Printf("Barang '%s' dengan harga %.2f berhasil ditambahkan.\n", namaBarang, hargaBarang)

		case "2":
			// Selesaikan transaksi
			fmt.Printf("\nTotal belanja: %.2f\n", totalBelanja)
			fmt.Println("Transaksi selesai. Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}