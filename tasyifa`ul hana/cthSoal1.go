package main
import "fmt"
func main() {
	var kata string
	var jumlah int

	fmt.Scan(&kata)
	fmt.Scan(&jumlah)

	for i := 1; i <= jumlah; i++ {
		fmt.Println(kata)
	}
}