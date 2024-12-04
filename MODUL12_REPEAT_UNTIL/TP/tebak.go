package main
import "fmt"

func main() {
    secretNumber := 5
    var guess int
    for {
        fmt.Print("Tebak angka rahasia antara 1 hingga 10: ")
        fmt.Scan(&guess)
        if guess == secretNumber {
            fmt.Println("Selamat, tebakanmu benar!")
            break 
        } else {
            fmt.Println("Tebakanmu salah, coba lagi.")
        }
    }
}