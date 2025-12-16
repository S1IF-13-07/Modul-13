package main
import "fmt"

func main() {
	var a, digit int
	fmt.Print("Masukkan sebuah angka bilangan bulat : ")
	fmt.Scan(&a)

	for a > 0 {
		a = a / 10
		digit++
	}
	fmt.Printf("Jumlah digit dari bilangan tersebut adalah: %d\n", digit)
}