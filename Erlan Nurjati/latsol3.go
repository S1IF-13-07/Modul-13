package main

import "fmt"

func main() {
	var target int
	fmt.Print("Masukan target: ")
	fmt.Scan(&target)


	total := 0
	jumlahDonatur := 0

	for selesai := false; !selesai; {
		var donasi int
		fmt.Print("Masukan donasi: ")
		fmt.Scan(&donasi)
		jumlahDonatur++
		total += donasi

		fmt.Printf(
			"Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
			jumlahDonatur, donasi, total,
		)

		selesai = total >= target
	}

	fmt.Printf(
		"Target tercapai! Total donasi: %d dari %d donatur.\n",
		total, jumlahDonatur,
	)
}
