package main  
import "fmt"  
func main() {  
    var umur int  
    var kk bool  
    fmt.Scan(&umur, &kk)  
    if umur >= 17 && kk {  
        fmt.Println("bisa buat id card")  
    }else{  
        fmt.Println("tidak bisa membuat id card")  
    } 
} 