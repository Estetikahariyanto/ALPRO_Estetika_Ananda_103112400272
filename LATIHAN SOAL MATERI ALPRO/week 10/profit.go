package main

import "fmt"

func main() {
	var lastMonth, thisMonth float64
	fmt.Print("Masukkan keuntungan bulan sebelumnya dan bulan ini (pisahkan dengan spasi): ")
	fmt.Scan(&lastMonth, &thisMonth)

	switch {
	case thisMonth > lastMonth:
		fmt.Printf("Peningkatan sebesar %.2f\n", thisMonth-lastMonth)
	case thisMonth < lastMonth:
		fmt.Printf("Penurunan sebesar %.2f\n", lastMonth-thisMonth)
	default:
		fmt.Println("Tetap")
	}
}