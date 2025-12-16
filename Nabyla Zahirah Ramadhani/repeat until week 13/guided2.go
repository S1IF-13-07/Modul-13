package main
import "fmt"
func main (){
	var n int
	var continueLoop bool
	for continueLoop= true; continueLoop; {
		fmt.Scan(&n)
		continueLoop = n <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", n)
}