package main

import "fmt"

func main(){
    var angka int

    fmt.Scanln(&angka)

	digit1 := angka / 100   
    digit2 := (angka / 10) % 10
    digit3 := angka % 10

	fmt.Println(digit1 > digit2 && digit2 > digit3)
}

