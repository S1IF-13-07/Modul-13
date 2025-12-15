package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan y: ")
	fmt.Scan(&y)

	hasil := 0
	temp := x

	for temp >= y {
		temp = temp - y
		hasil++
	}

	fmt.Printf("Hasil %d div %d = %d\n", x, y, hasil)
}
