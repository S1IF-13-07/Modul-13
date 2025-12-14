package main

import "fmt"

func main() {
	var angka int
	var lanjut bool

	for lanjut = true; lanjut; {
		fmt.Scan(&angka)
		lanjut = angka <= 0
	}

	fmt.Printf("%d adalah bilangan bulat positif\n", angka)
}
