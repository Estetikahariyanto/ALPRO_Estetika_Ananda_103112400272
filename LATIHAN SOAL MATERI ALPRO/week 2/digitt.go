package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)

    // x kurang atau = 999
    if x >= 0 && x <= 999 {
        d1 := x / 100         // Digit pertama (ratusan)
        d2 := (x / 10) % 10   //  kedua (puluhan)
        d3 := x % 10          //  ketiga (satuan)

        fmt.Println(d1, d2, d3)
    } else {
        fmt.Println("Bilangan antara 0 dan 999")
    }
}
