package main
import "fmt"

func main() {
	var bil, digit int
	digit = 0
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bil)

	for bil > 0 {
		bil = bil / 10
		digit++
	}
	fmt.Print("jumalh digit: ", digit)
}