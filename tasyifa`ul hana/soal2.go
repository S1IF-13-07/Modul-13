package main
import (
	"fmt"
	"math"
)
func main() {
	var angka float64

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&angka)

	bulatAtas := math.Ceil(angka)

	for i := angka + 0.1; i <= bulatAtas; i += 0.1 {
		if i == math.Floor(i) {
			fmt.Printf("%.0f\n", i)
		} else {
			fmt.Printf("%.1f\n", i)
		}
	}
}