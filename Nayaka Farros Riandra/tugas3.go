package main

import "fmt"

func main(){
	var target, total, donatur int
	fmt.Println("masukkan target donasi : ")
	fmt.Scan(&target)

	total = 0
	donatur = 0

	fmt.Println("masukkan jumlah uang untuk didonasikan : ")

	for total < target {
		var donasi int
		fmt.Scan(&donasi)
		donatur+=1
		total += donasi

		fmt.Printf("Donatur %d: menyumbang %d. total terkumpul: %d\n", donatur, donasi, total)
	}

	fmt.Printf("target tercapai total donasi: %d dari %d donatur. \n", total, donatur)
}