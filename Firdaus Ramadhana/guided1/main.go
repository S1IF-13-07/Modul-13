package main 
import "fmt"

func main() {
	var teks string
	var cetak int

	fmt.Print("Masukkan teks dan berapakali akan dicetak: ")
	fmt.Scanln(&teks, &cetak)

	for i := 0; i < cetak; i++ {
		fmt.Println(teks)
	}
}