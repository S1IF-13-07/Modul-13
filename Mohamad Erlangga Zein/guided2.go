package main 
import "fmt"

func main() {
	var n int
	var aselole bool

	for aselole = false; !aselole; {
		fmt.Print("masukkan angka: ")
		fmt.Scan(&n)
		
		aselole = (n * -1) < 0
	}
	fmt.Println(n, "adalah bilangan positif")
}