package main
import "fmt"
func main() {
	var x int
	var kondisi bool

	for kondisi = true; kondisi;{
		fmt.Scan(&x)
		kondisi = x <= 0 
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", x)
}