package main

import "fmt"

func main(){
	var num int
	for kondisi := false; !kondisi;{
		fmt.Scan(&num)
		if num > 0 {
			fmt.Print(num, " Adalah bilangan bulat positif")
			kondisi = true
		}
	}

}