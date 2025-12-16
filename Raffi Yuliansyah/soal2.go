package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)

	batas := math.Ceil(n)

	saatIni := int(n * 10)
	nLanjutan := int(batas * 10)

	for i := saatIni + 1; i <= nLanjutan; i++ {
		if i%10 == 0 {
			fmt.Println(i / 10)
		} else {
			fmt.Printf("%.1f\n", float64(i)/10.0)
		}
	}
}
