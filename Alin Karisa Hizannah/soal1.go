package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sisa := n
	jumlahDigit := 0

	for selesai := false; !selesai; {
		sisa = sisa / 10
		jumlahDigit++
		selesai = (sisa == 0)
	}

	fmt.Println(jumlahDigit)
}
