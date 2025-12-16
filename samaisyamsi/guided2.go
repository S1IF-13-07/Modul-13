package main
import "fmt"
func main() {
	
    var n int

    for {
        fmt.Print("Masukkan bilangan bulat positif: ")
        fmt.Scan(&n)

        if n > 0 {
            fmt.Printf("%d adalah bilangan bulat positif\n", n)
            break
        } else {
            fmt.Println("Input bukan bilangan bulat positif, coba lagi.")
        }
    }
}

