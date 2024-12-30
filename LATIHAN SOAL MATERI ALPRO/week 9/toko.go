package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(a+a*5/100, b+b*5/100, c+c*5/100)
}