package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Print("Masukkan sebuah angka desimal : ")
	fmt.Scan(&a)

	batas := math.Ceil(a)
	saatini := int(a * 10)
	nLanjutan := int(batas * 10)

	for i := saatini; i <= nLanjutan; i++ {
		fmt.Printf("%.1f ", float64(i)/10)
	}
	fmt.Println()

}
