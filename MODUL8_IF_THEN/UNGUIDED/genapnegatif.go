package main
import(
	"fmt"
)

func main() {
	var bilangan int
	fmt.Print("masukan bilangan")
	fmt.Scan(&bilangan)

	// mengecek bilangan genap dan negatif
	if bilangan < 0 && bilangan%2 == 0 {
		fmt.Println("genap negatif")
	} else {
		fmt.Println("bukan")
	}
}

