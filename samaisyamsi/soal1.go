package main
import "fmt"
func main() {
	var n int
	fmt.Print("Masukan biangan")
	fmt.Scan(&n)

	 jumlahdigit := 0

	 for {
		jumlahdigit++
		n = n / 10

		if n == 0 {
			break 
		}

	 }
	 fmt.Println(jumlahdigit)
}