package main
import "fmt"
func main() {
	var target, donasi int

	fmt.Print("masukan target: ")
	fmt.Scan(&target)

	total := 0
	donatur := 0 

	for {
		fmt.Scan(&donasi)
		donatur++
		total = total + donasi
		fmt.Printf("Donatur %d : Menyumbang %d. Total terkumpul: %d\n", donatur, donasi, total)
		
		if total >= target {
			break
		}
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, donatur)
}