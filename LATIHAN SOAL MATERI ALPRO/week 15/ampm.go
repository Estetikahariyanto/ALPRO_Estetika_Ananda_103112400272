package main
import "fmt"
func main() {
    var jam24, jam12 int

    //input jam dalam format 24 jam
    fmt.Print("Masukkan jam dalam format 24 jam: ")
    fmt.Scan(&jam24)

    //konversikan ke format 12 jam
    if jam24%12 == 0 {
        jam12 = 12
    } else {
        jam12 = jam24 % 12
    }

    //tentukan AM/PM
    if jam24 < 12 {
        fmt.Printf("%02d AM\n", jam12)
    } else {
        fmt.Printf("%02d PM\n", jam12)
    }
}