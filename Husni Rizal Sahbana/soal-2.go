package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	batas := int(x)
	if float64(batas) < x {
		batas++
	}

	for x < float64(batas) {
		x += 0.1

		if x < float64(batas) {
			fmt.Printf("%.1f\n", x)
		} else {
			fmt.Println(batas)
		}
	}
}