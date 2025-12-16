package main

import (
	"fmt"
	"math"
)

func main(){
	var n float64
	fmt.Scan(&n)

	batas := math.Ceil(n)

	saatini := int(n*10)
	nlanjutan := int(batas*10)

	for i := saatini + 1; i <= nlanjutan; i++ {
		if i%10 == 0 {
			fmt.Println(i/10)
		} else {
			fmt.Println("", float64(i) /10.0)
		}
	}
}