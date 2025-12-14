package main

import "fmt"

func main() {
	var benda string
	var jumlah int

	fmt.Scan(&benda, &jumlah)

	counter := 0

	for done := false; !done; {
		fmt.Println(benda)
		counter++
		done = (counter >= jumlah)
	}
}
