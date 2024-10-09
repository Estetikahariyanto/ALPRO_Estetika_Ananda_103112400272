package main

import "fmt"

func main() {
    var bmi, height, weight float64

    // Input BMI dan tinggi badan dalam meter
    fmt.Print("Masukkan nilai BMI: ")
    fmt.Scan(&bmi)
    fmt.Print("Masukkan tinggi badan (dalam meter): ")
    fmt.Scan(&height)

    // Rumus: Berat badan (kg) = BMI * Tinggi^2
    weight = bmi * (height * height)

    // Output berat badan dalam kilogram
    fmt.Printf("Berat badan Anda adalah %.2f kilogram\n", weight)
}
