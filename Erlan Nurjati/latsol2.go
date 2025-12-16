package main

import "fmt"

func main() {
	var n float64
	fmt.Print("Input: ")
	fmt.Scan(&n)

	awal := int(n * 10)
	batas := int(n)

	if n > float64(batas) {
		batas += 1
	}

	for i := awal + 1; i <= batas*10; i++ {
		if i%10 == 0 {
			fmt.Println(i / 10)
		} else {
			fmt.Printf("%.1f\n", float64(i)/10)
		}
	}
}