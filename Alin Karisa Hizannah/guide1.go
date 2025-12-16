package main

import "fmt"

func main() {
	var kata string
	var jumlahKata int
	
	fmt.Print("Masukkan kata dan jumlah kata: ")
	fmt.Scan(&kata, &jumlahKata)
	
	i := 0
	for kondisi := false; !kondisi; {
		fmt.Println(kata)
		i++
		kondisi = (i >= jumlahKata)
	}
}
