package main

import "fmt"

func main(){
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	for kondisi := false; !kondisi;{
		a=a-b
		fmt.Println(a)

		if a<=0{
			kondisi= a == 0
			fmt.Println(kondisi)

		kondisi=true
		}
		
	}
}