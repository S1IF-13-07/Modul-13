package main 
import (
	"fmt"
	"math"
)
func main() {

	var bildesimal float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&bildesimal)
	batas := math.Ceil(bildesimal)
	bilsaatini := int(bildesimal * 10)
	bilLanjutan := int(batas * 10)
	for i := bilsaatini + 1; i <= bilLanjutan; i++{
		if i%10 == 0 {
			fmt.Println(i / 10)
		} else {
			fmt.Printf("%.1f\n", float64(i) / 10.0)
		}
	}
}