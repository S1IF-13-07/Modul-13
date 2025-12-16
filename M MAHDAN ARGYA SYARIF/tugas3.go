package main 
import "fmt"

func main() {
	var target, total, jumlahDonatur int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	total = 0
	jumlahDonatur = 0

	fmt.Println("Masukkan jumlah uang untuk di donasikan: ")

	for total <= target {
		var donasi int
		fmt.Scan(&donasi)
		jumlahDonatur += 1

		total += donasi

		fmt.Printf("Donatur %d: menyumbang %d. Total terkumpul: %d\n ", jumlahDonatur, donasi, total)
	}
	fmt.Printf("Target tercapai! total donasi: %d dari %d donatur. \n", total, jumlahDonatur)
}