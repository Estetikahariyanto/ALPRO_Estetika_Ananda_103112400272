package main
import "fmt"

func main() {
	var n, k, j int
	fmt.Scan(&n)
	for j = 1; j <= n; j++ {
	for k = 1; k <= n; k++ {
		fmt.Print(j, " ")
	}
	fmt.Println()
	}
}