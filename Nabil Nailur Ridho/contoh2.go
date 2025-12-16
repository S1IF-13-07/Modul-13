package main

import "fmt"

func main(){
	var a int
	var kondisi bool	
		for kondisi = false; !kondisi;{
			fmt.Print("Masukkan Angka: ")
			fmt.Scan(&a)

			kondisi = (a * -2) < 0
		}
		fmt.Println(a, "adalah bilangan bulat")
}