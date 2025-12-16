package main

import "fmt"

func main(){
	var kata string
	var kondisi bool
	var jumlahKata int

	fmt.Print("Masukkan Kata dan Jumlah Kata: ")
	fmt.Scan(&kata, &jumlahKata)

	for kondisi = false; !kondisi; {
		fmt.Println(kata)
		jumlahKata--
		kondisi = jumlahKata == 0
	}
}