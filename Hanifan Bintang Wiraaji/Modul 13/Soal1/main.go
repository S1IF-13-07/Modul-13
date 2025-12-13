package main

import "fmt"

func main() {
	var a int
	i := 0
	fmt.Scan(&a)
	for a > 0 {
		i++
		a /= 10
	}
	fmt.Print(i)
}