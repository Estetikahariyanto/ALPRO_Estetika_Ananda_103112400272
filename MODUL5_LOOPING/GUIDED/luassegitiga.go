package main
import "fmt"
func main() {
 var j, alas, tinggi, n int
 var area float64
 fmt.Scan(&n)
 for j = 1; j <=n; j+=1 {
 fmt.Scan(&alas, &tinggi)
 area = 0.5 * float64(alas * tinggi)
 fmt.Println(area)
 }
}