package main

import "fmt"

func main() {
    var x, y int
    isKelipatan := true

    fmt.Scan(&x, &y)

    for {
        x -= y
        fmt.Println(x)

        if x < 0 {
            isKelipatan = false
            break
        }

        if x == 0 {
            break
        }
    }

    fmt.Println(isKelipatan)
}