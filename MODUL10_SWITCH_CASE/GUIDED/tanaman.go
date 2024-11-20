package main

import "fmt"

func main() {
	var namaTanaman string
	fmt.Scanln(&namaTanaman) // Menggunakan Scanln agar bisa menangani nama tanaman dengan spasi

	switch namaTanaman {
	case "nepenthes", "drosera":
		fmt.Println("Termasuk Tumbuhan Karnivora.")
		fmt.Println("Asli Indonesia.")
	case "venus", "sarracenia":
		fmt.Println("Termasuk Tumbuhan Karnivora.")
		fmt.Println("Bukan Asli Indonesia.")
	default:
		fmt.Println("Bukan Tumbuhan Karnivora.")
	}
}