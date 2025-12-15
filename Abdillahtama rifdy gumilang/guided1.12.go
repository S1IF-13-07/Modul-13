package main

import "fmt"

func main() {
	var n, j int
	fmt.Scan(&n)
	j = n
	for j > 1 {
		fmt.Println(j, "x")
		j = j - 1
	}
	fmt.Println(1)
}
