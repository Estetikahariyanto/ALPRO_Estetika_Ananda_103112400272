package main

import "fmt"

func main() {
	var base, exponent int
	fmt.Scan(&base, &exponent)

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}

	fmt.Println(result)
}
