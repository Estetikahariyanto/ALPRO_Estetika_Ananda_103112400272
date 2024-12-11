package main

import "fmt"
func main() {
    var n, count int
    fmt.Print("Masukkan bilangan bulat positif n: ")
    fmt.Scan(&n)
    if n <= 0 {
        fmt.Println("Masukkan bilangan positif!")
        return
    }
    for i := 1; i <= n; i++ {
        if i%2 != 0 {
            count++
        }
    }
    fmt.Printf("Banyaknya bilangan ganjil dari 1 hingga %d adalah %d.\n", n, count)
}