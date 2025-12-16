package main 
import (
	"fmt"
	"math"
)

func main() {
	var bil float64
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bil)

	batasAtas := math.Ceil(bil)

	nilaiAwal := int(bil * 10)
	nilaiLanjutan := int(batasAtas * 10)
	

	for i := nilaiAwal + 1; i <= nilaiLanjutan; i++{
		if i%10 == 0 {
			fmt.Println(i / 10)
		} else {
			fmt.Printf("%.1f\n", float64(i) / 10.0)
		}
	}
}