package main

import "fmt"

func main() {
	var target int
	var donasi int
	var total int
	var counter int
	fmt.Scan(&target)
	total = 0
	counter = 0
	for selesai := false; !selesai; {
		fmt.Scan(&donasi)
		counter++
		total = total + donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", counter, donasi, total)
		selesai = total >= target
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, counter)
}