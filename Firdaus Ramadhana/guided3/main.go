package main 
import "fmt"

func main() {
	var x, y int 
	var kondisi bool
	fmt.Print("Masukkan dua angka: ")
	fmt.Scanln(&x, &y)

	for kondisi = false; !kondisi; {
		x -= y
		fmt.Println(x)
		if x <= 0 {
			kondisi = x == 0
			fmt.Println(kondisi)
			kondisi = true
		}
	}
}