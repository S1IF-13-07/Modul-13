package main

import "fmt"

func main() {
	var n, jumlah int
	fmt.Scan(&n)
	jumlah = 0
	for {
		n = n / 10
		jumlah++
		if n == 0 {
			break
		}
	}
	fmt.Printf("%d\n", jumlah)
}
