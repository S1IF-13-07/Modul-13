package main

import "fmt"

func main() {
	var a string
	var b int
	fmt.Print("Masukan Kamu :")
	fmt.Scan(&a, &b)
	counter := 0
	for done := false; !done; {
		fmt.Println(a)
		counter++
		done = (counter >= b)
	}
}
