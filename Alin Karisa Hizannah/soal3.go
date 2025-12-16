package main

import "fmt"

func main() {
	var target int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	totalDonasi := 0
	jumlahDonatur := 0

	for selesai := false; !selesai; {
		var donasi int
		fmt.Print("Masukkan jumlah donasi: ")
		fmt.Scan(&donasi)

		jumlahDonatur++
		totalDonasi = totalDonasi + donasi

		fmt.Printf(
			"Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
			jumlahDonatur, donasi, totalDonasi,
		)

		selesai = (totalDonasi >= target)
	}

	fmt.Printf(
		"Target tercapai! Total donasi: %d dari %d donatur.\n",
		totalDonasi, jumlahDonatur,
	)
}
