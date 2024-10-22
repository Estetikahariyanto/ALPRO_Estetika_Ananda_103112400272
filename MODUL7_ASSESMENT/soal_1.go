package main

import "fmt"

func main() {
    var N int

    // Meminta input dari pengguna
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&N)

    // Mencetak hasil kuadrat dari 1 hingga N
    for i := 1; i <= N; i++ {
        fmt.Printf("%d^2 = %d\n", i, i*i)
    }
}
