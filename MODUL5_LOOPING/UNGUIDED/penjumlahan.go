package main

import "fmt"

// Fungsi untuk menjumlahkan sekumpulan bilangan
func sum(numbers []int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    var n int
    fmt.Print("Masukkan jumlah bilangan yang ingin dijumlahkan: ")
    fmt.Scan(&n)
    
    numbers := make([]int, n)
    fmt.Println("Masukkan bilangan-bilangannya:")
    for i := 0; i < n; i++ {
        fmt.Printf("Bilangan %d: ", i+1)
        fmt.Scan(&numbers[i])
    }
    
    hasil := sum(numbers)
    fmt.Printf("Hasil penjumlahan bilangan-bilangan tersebut adalah: %d\n", hasil)
}
