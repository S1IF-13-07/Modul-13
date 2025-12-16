package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukan n: ")
	fmt.Scan(&n)

	count := 0
	for selesai := false; !selesai; {
		n = n / 10
		count++
		selesai = n == 0
	}

	fmt.Println(count)
}
