package main
import "fmt"
func main() {
	
	var number, d, counter int
	fmt.Scan(&number)
	counter = 0
	for number > 0 {
		d = number % 10
		if d%2 == 0 && d != 0 {
			counter = counter + 1
		}
		number = number / 10
	}
	fmt.Print(counter)
}