package main

import "fmt"

func main() {
	var a, i int
	fmt.Scan(&a)
	b := 0
	for b < a {
		c := 0
		fmt.Scan(&c)
		b += c
		i++
		fmt.Printf("Donatur %v: Menyumbang %v. Total terkumpul: %v\n", i, c, b)
	}
	fmt.Print("Target tercapai! Total donasi: ", b, " dari ", i, " donatur")
}
