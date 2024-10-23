package main

import "fmt"

// Fungsi untuk menghitung hasil perkalian tanpa menggunakan operator *
func multiply(a int, b int) int {
    result := 0
    // Jika b negatif, konversi ke positif dan gunakan penanda negatif
    negative := false
    if b < 0 {
        negative = true
        b = -b
    }
    // Tambahkan a ke result sebanyak b kali
    for i := 0; i < b; i++ {
        result += a
    }
    // Jika b awalnya negatif, hasil juga harus negatif
    if negative {
        result = -result
    }
    return result
}

func main() {
    var a, b int
    fmt.Print("Masukkan angka pertama: ")
    fmt.Scan(&a)
    fmt.Print("Masukkan angka kedua: ")
    fmt.Scan(&b)
    
    hasil := multiply(a, b)
    fmt.Printf("Hasil perkalian %d dan %d adalah: %d\n", a, b, hasil)
}
