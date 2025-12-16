package main
import "fmt"
func main() {
    var x float64
    var target float64
    fmt.Scan(&x)
    n := int(x)
    if float64(n) < x {
        target = float64(n + 1)
    } else {
        target = float64(n)
    }
    for done := false; !done; {
        x = x + 0.1
        fmt.Printf("%.1f\n", x)
        done = x >= target
    }
}