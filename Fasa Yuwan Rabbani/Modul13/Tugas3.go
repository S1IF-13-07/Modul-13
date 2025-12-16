package main
import "fmt"

func main() {
	var a, t, d int
	fmt.Print("Masukkan sebuah angka : ")
	fmt.Scan(&a)

	t = 0
	d = 0
	
	fmt.Println("Masukkan nominal untuk didonasikan : ")

	for t <= a {
		var donasi int
		fmt.Scan(&donasi)
		d += 1

		t += donasi
		fmt.Printf("Donatur %d: menyumbang : %d. Total Terkumpul : %d\n", d, donasi, t)
	}
	fmt.Printf("Target Tercapai! total donasi %d dari %d Donatur. \n", t, d)

}