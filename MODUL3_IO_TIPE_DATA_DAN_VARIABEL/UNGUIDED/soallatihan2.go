package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi konstanta Pi
	const pi = 3.1415926535

	// Input dari user
	var r int
	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	// Menghitung volume bola
	volumeBola := (4.0 / 3.0) * pi * math.Pow(float64(r), 3)

	// Menghitung luas permukaan bola
	luasBola := 4 * pi * math.Pow(float64(r), 2)

	// Menampilkan hasil
	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volumeBola, luasBola)
}
