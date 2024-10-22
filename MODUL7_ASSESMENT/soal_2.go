package main

import "fmt"

func main() {
    var jumlahBarang int
    var totalPoin int

    // Meminta input jumlah barang dari pengguna
    fmt.Print("Masukkan jumlah barang yang dibeli: ")
    fmt.Scan(&jumlahBarang)

    // Menghitung poin berdasarkan jumlah barang yang dibeli
    if jumlahBarang > 5 {
        totalPoin = (5 * 10) + ((jumlahBarang - 5) * 15) // 10 poin untuk 5 barang pertama, 15 poin untuk barang ke-6 dan seterusnya
    } else {
        totalPoin = jumlahBarang * 10 // 10 poin untuk setiap barang jika <= 5 barang
    }

    // Menampilkan total poin yang didapatkan
    fmt.Printf("Total poin yang didapatkan: %d\n", totalPoin)
}