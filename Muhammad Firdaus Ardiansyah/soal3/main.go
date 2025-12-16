package main 
import "fmt"

func main() {
	var target, total, donatur int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	total = 0
	donatur = 0

	fmt.Println("Masukkan jumlah uang untuk di donasikan: ")

	for total <= target {
		var donasi int
		fmt.Scan(&donasi)
		donatur += 1

		total += donasi

		fmt.Printf("Donatur %d: menyumbang %d. Total terkumpul: %d\n ", donatur, donasi, total)
	}
	fmt.Printf("Target tercapai! total donasi: %d dari %d donatur. \n", total, donatur)
}