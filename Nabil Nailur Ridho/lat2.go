package main

import "fmt"

func main() {
	var n float64
	fmt.Print("Masukkan sebuah bilangan desimal: ")
	fmt.Scan(&n)

	bilBul := int(n)
	target := bilBul

	if n != float64(bilBul) {
		target = bilBul + 1
	}
	for kondisi:= false; !kondisi; {
		n = n + 0.1
		fmt.Printf("%.1f\n", n)

		kondisi = n >= float64(target) - 0.000001
	}
}