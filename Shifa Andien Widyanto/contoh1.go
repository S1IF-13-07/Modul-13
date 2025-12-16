package main 
import "fmt" 
func main() {
	var kata string
	var x int
	fmt.Print("Masukan suatu kata dan jumlah perulangan : ")
	fmt.Scan(&kata, &x)
	counter := 0
	for done := false; !done; {
		fmt.Println(kata)
		counter++
		done = (counter >= x)
	}
}