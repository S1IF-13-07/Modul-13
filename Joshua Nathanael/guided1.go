package main

import "fmt"

func main() {
	var kata string
	var jumlah int
	
	fmt.Print("Masukkan sebuah kata & jumlahnya: ")
	fmt.Scan(&kata, &jumlah)

	for kondisi := false; !kondisi; {
		fmt.Println(kata)
		jumlah--
		kondisi = jumlah == 0
	}
}