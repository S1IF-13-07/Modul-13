package main

import "fmt"

func main() {
	var n int
	var count int
	fmt.Scan(&n)
	count = 0
	for done := false; !done; {
		n = n / 10
		count++
		done = n == 0
	}
	fmt.Println(count)
}