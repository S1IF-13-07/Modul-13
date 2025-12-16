package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	limit := math.Ceil(x)

	for x < limit {
		x += 0.1

		x = math.Round(x*10) / 10

		if x == limit {
			fmt.Printf("%.0f\n", x)
		} else {
			fmt.Printf("%.1f\n", x)
		}
	}
}