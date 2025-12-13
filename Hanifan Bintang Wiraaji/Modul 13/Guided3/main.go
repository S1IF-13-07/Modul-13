package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for kondisi := false; !kondisi; {
		a = a - b
		fmt.Println(a)
		kondisi = a <= 0
	}
	fmt.Print(a == 0)
}
