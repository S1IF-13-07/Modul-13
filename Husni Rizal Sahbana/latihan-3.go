package main

import "fmt"

func main() {
    var a, b int
    var kondisi bool
    fmt.Scan(&a, &b)
    for kondisi = false; !kondisi; {
        a = a - b
        fmt.Println(a)
        kondisi = a <= 0
    }
    fmt.Println(a == 0)
}
