package main
import "fmt"
func main() {
	var angka int
	jumlahDigit := 0

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&angka)

	for angka > 0 {
		angka = angka / 10
		jumlahDigit++
	}

	fmt.Println("Jumlah digit:", jumlahDigit)
}