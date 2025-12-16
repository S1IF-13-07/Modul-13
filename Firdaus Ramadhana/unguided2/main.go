package main

import "fmt"

func main() {
	var a float64
	fmt.Print("Masukkan angka desimal: ")
	fmt.Scan(&a)

	b := int(a * 10)
	limit := ((b / 10) + 1) * 10

	for b < limit {
		b++
		fmt.Printf("%.1f\n", float64(b)/10)
	}
}
