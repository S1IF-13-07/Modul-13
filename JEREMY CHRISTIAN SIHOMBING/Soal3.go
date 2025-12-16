package main 
import "fmt"
func main() {

	var targetdonasi, total, donatur int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&targetdonasi)
	total = 0
	donatur = 0
	fmt.Println("Masukkan jumlah uang untuk didonasikan: ")
	for total <= targetdonasi {
		var donasi int
		fmt.Scan(&donasi)
		donatur += 1
		total += donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n ", donatur, donasi, total)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur. \n", total, donatur)
}