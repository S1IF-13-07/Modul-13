package main
import "fmt"
func main() {
	var a float64
	
	fmt.Print("Masukan bilangan desimal: ")
	fmt.Scan(&a)
	nilai := int(a * 10)
	batas := (int(a) + 1) * 10

	for { 
	nilai = nilai + 1
	if nilai == batas {
		fmt.Println(batas / 10)
		break
	} else {
		fmt.Printf("%.1f\n", float64(nilai)/10)
		}
	}
}