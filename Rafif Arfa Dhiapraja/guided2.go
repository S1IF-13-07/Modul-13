package main

import "fmt"

func main() {
	var a int
	var n bool

	for n = false; !n; {
		fmt.Print("Masukkan sebuah angka: ")
		fmt.Scan(&a)

		n = (a * -2) < 0
	}

	fmt.Println(n, "adalah angka positif")
}