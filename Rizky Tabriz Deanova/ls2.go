package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Scan(&input)
	target := math.Ceil(input)
	for {
		input = input + 0.1
		input = math.Round(input*10) / 10
		if input == float64(int(input)) {
			fmt.Printf("%.0f\n", input)
		} else {
			fmt.Printf("%.1f\n", input)
		}
		if input >= target {
			break
		}
	}
}
