package main
import "fmt"
func main() {
	var n int
	jumlahDigit := 0

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	for n > 0 {
		jumlahDigit++
		n = n / 10
	}
	fmt.Println("Jumlah digit:", jumlahDigit)
}
