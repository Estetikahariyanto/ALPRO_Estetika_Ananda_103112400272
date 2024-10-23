package main
import (
	"fmt"
)

func main() {
	var qirat int
// minta input dari user
	fmt.Print("masukan jumlah uang dalam qirat: ")
	fmt.Scan (&qirat)

// konversi qirat ke fals, dirham dan dinar
fals := qirat / 6
sisaQirat := qirat % 6

dirham := fals / 10
sisaFals := fals % 10

dinar := dirham / 10
sisaDirham := dirham % 10

// menampilkan hasil
fmt.Printf("Hasil konversi:\n")
fmt.Printf("%d dinar, %d dirham, %d fals, %d qirat\n", dinar, sisaDirham, sisaFals, sisaQirat)
}