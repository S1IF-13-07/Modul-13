package main

import "fmt"

func main() {
	var a int
	for kondisi := false; !kondisi; {
		fmt.Scan(&a)
		kondisi = a > 0
	}
	fmt.Print(a, " Bilangan bulat positif")
}
