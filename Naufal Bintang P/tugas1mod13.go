package main

import "fmt"

func main() {
    var n int
    var digit int

    fmt.Scan(&n)

    for n != 0 {
        n = n / 10
        digit++
    }

    fmt.Println(digit)
}