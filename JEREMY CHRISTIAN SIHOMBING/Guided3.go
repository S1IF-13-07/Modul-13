package main  
import "fmt"  
func main() { 

    var bilx, bily int
    var kondisi bool 
    fmt.Scan(&bilx, &bily) 
    for kondisi = false; !kondisi; { 
        bilx = bilx - bily 
        fmt.Println(bilx) 
        kondisi = bilx <= 0 
    } 
    fmt.Println(bilx == 0) 
}