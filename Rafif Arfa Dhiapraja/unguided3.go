package main

import "fmt"

func main() {
	var target, jumlah, donasi int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)
	
	jumlah = 0
	donatur := 0

	for kondisi := false; !kondisi; {
		fmt.Print("Masukkan donasi: ")
		fmt.Scan(&donasi)
		
		jumlah += donasi
		donatur++

		fmt.Printf("Donatur %d: Menyumbang %d. Total Terkumpul: %d\n", donatur, donasi, jumlah)

		kondisi = jumlah >= target

	}

	fmt.Printf("Target tercapai total donasi: %d dari %d donatur\n", jumlah, donatur)
}