package main
import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	for kondisi := false; !kondisi; {
		x = x - y
		fmt.Println(x)
		if x <= 0{
			kondisi = x == 0
			fmt.Print(kondisi)

			kondisi = true
		}
	}
}