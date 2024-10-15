package main

import "fmt"

func main() {
    // Meminta input jumlah baris dari pengguna
    var n int
    fmt.Print("Masukkan jumlah baris: ")
    fmt.Scanln(&n)

    // Mencetak segitiga bintang
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}