package main

import "fmt"

func main() {
    var nama string
    var beratBadan, tinggiBadan float64

    // Input nama, berat badan, tinggi badan
    fmt.Scan(&nama, &beratBadan, &tinggiBadan)

    // Hitung BMI
    bmi := beratBadan / (tinggiBadan * tinggiBadan)

    // hasilnya
    fmt.Printf("Nama: %s, Berat Badan: %.2f kg, Tinggi Badan: %.2f m, BMI: %.2f\n", nama, beratBadan, tinggiBadan, bmi)
}
