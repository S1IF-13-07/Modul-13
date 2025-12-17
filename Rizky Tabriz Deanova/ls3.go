package main

import "fmt"

func main() {
	var target, donasi, total, donatur int
	fmt.Scan(&target)
	total = 0
	donatur = 0
	for {
		fmt.Scan(&donasi)
		total = total + donasi
		donatur = donatur + 1
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donatur, donasi, total)
		if total >= target {
			fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, donatur)
			break
		}
	}
}
