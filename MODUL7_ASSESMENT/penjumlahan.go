package main

import (
	"fmt"
)

func main () {
	var x, y int
// meminta input dari user
fmt.Print ("masukan nilai x: ")
fmt.Scan(&x)
fmt.Print ("masukan nilai y: ")
fmt.Scan(&y)
// validasi input x<=y
	if x > y {
		fmt.Println ("nilai x harus lebih kecil atau sama dengan y.")
		return
	}
// menjumlahkan nilai x sampai y
total :=0
for i:= x; i<= y; i++ {
	total += i
}
// menampilkan hasil 
fmt.Printf ("jumlah bilangan dari %d sampai %d adalah %d\n", x, y, total)
}
