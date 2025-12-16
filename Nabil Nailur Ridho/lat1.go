package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	jumlahdigit := 0

	for kondisi := false; !kondisi; {
		bilangan = bilangan / 10
		jumlahdigit++
		kondisi = bilangan == 0
	}
	fmt.Println("Jumlah Digit adalah:", jumlahdigit)

}