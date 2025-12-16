package main

import "fmt"

func main() {
	var target int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	total := 0
	donatur := 0

	for total <= target {
		var donasi int
		fmt.Print("Masukkan jumlah donasi: ")
		fmt.Scan(&donasi)

		donatur++
		total += donasi

		fmt.Printf(
			"Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
			donatur, donasi, total,
		)
	}

	fmt.Printf(
		"Target tercapai! Total donasi: %d dari %d donatur.\n",
		total, donatur,
	)
}
