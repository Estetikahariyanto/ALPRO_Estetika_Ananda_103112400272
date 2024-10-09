package main
import "fmt"
func main() {
 var beratbadan, tinggibadan, bmi float64
 fmt.Scan(&beratbadan, &tinggibadan) 
 bmi = beratbadan / (tinggibadan * tinggibadan) 
 fmt.Printf("%.2f", bmi) 
}
