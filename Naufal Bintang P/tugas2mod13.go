package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	start := int(x * 10)
	end := (start/10 + 1) * 10

	for i := start + 1; i <= end; i++ {
		if i == end {
			fmt.Println(i / 10)
		} else {
			fmt.Printf("%.1f\n", float64(i)/10)
		}
	}
}