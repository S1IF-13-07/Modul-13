package main
import "fmt"

func main() {
	var n, digit int
	digit = 0
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	for n > 0 {
		n = n / 10
		digit++
	}
	fmt.Print(digit)
}