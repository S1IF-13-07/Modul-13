package main

import "fmt"

func main(){
	var num int
	var word string
	var kondisi bool
	fmt.Scan(&word)
	fmt.Scan(&num)
	for kondisi = false; !kondisi; {
		fmt.Println(word)
		num--
		kondisi = num == 0
	}
}