package main

import "fmt"

func main() {
    const attempts = 5
    correctSequence := [4]string{"merah", "kuning", "hijau", "ungu"}
    isAllCorrect := true

    // Loop untuk 5 kali percobaan
    for attempt := 1; attempt <= attempts; attempt++ {
        fmt.Printf("Percobaan %d:\n", attempt)
        var sequence [4]string
        isCorrect := true

        // Memasukkan warna untuk 4 gelas reaksi
        for i := 0; i < 4; i++ {
            fmt.Printf("Masukkan warna untuk gelas ke-%d: ", i+1)
            fmt.Scan(&sequence[i])
            if sequence[i] != correctSequence[i] {
                isCorrect = false
            }
        }

        // Jika salah satu percobaan salah, set isAllCorrect menjadi false
        if !isCorrect {
            isAllCorrect = false
        }
    }

    // Menampilkan hasil akhir
    fmt.Println(isAllCorrect)
}