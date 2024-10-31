package main
	import "fmt"

func main() {
	var bilangan int
	var teks string
	fmt.Scan(&bilangan)
	teks = "tidak-positif"
	if bilangan > 0 {
		teks = "positif"
	}
fmt.Println(teks)
}

