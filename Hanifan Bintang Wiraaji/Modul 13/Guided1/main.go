package main

import (
	"fmt"
)

func main() {
	var kata string
	var a int
	fmt.Scan(&kata)
	fmt.Scan(&a)
	kondisi := 0

	for kondisi2 := false; !kondisi2; {
		fmt.Println(kata)
		kondisi++
		kondisi2 = kondisi == a
	}
}
