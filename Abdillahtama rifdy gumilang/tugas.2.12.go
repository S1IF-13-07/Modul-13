package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	fmt.Println("Digit dari kanan ke kiri:")
	for n > 0 {
		digit := n % 10
		fmt.Println(digit)
		n = n / 10
	}
}
