package main

import "fmt"

func main() {
	var nomor int
	var continueLoop bool
	for continueLoop = true; continueLoop; {
		fmt.Scan(&nomor)
		continueLoop = nomor <= 0
	}
	fmt.Printf("%d bilangan bulat positif\n", nomor)
}