package main

import (
    "fmt"
    "math"
)

// Fungsi untuk menghitung jarak antara dua titik
func distance(x1, y1, x2, y2 float64) float64 {
    return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
    var ax, ay, bx, by, cx, cy float64

    // Input koordinat titik A
    fmt.Print("Masukkan koordinat titik A (x y): ")
    fmt.Scan(&ax, &ay)

    // Input koordinat titik B
    fmt.Print("Masukkan koordinat titik B (x y): ")
    fmt.Scan(&bx, &by)

    // Input koordinat titik C
    fmt.Print("Masukkan koordinat titik C (x y): ")
    fmt.Scan(&cx, &cy)

    // Hitung panjang sisi AB, BC, dan CA
    ab := distance(ax, ay, bx, by)
    bc := distance(bx, by, cx, cy)
    ca := distance(cx, cy, ax, ay)

    // Tentukan sisi terpanjang
    longest := math.Max(ab, math.Max(bc, ca))

    // Tampilkan hasil dengan dua angka di belakang koma
    fmt.Printf("Sisi terpanjang dari segitiga adalah %.2f\n", longest)
}
