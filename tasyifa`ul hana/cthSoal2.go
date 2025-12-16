package main
import "fmt"
func main() {
	var angka int
	for {
		fmt.Print("Masukkan bilangan bulat positif: ")
		fmt.Scanln(&angka)

		if angka > 0 {
			break 
		}
	}

	fmt.Printf("%d adalah bilangan bulat positif\n", angka)
}