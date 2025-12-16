package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan float64
	fmt.Scan(&bilangan)

	batas := math.Ceil(bilangan)
	nilai := bilangan

	for selesai := false; !selesai; {
		nilai = nilai + 0.1

		nilai = math.Round(nilai*10) / 10

		fmt.Printf("%.1f\n", nilai)
		selesai = (nilai >= batas)
	}
}
