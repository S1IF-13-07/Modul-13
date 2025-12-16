package main
import "fmt"
func main() {
	var a int

	fmt.Print("Masukan bilangan bulat positif: ")
	fmt.Scan(&a)
	jumlah := 0
	for {
		jumlah++
		a = a / 10
		if a == 0 {
			break
		}
	}
	fmt.Println("Jumlah digit: ", jumlah)
}