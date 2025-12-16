package main
import "fmt"
func main() {
 var number int
 var continueLoop bool

 for continueLoop = true; continueLoop; {
 fmt.Scan(&number)
 continueLoop = number <= 0
 }
 fmt.Printf("%d ADALAH BILANGAN BULAT POSISTIF \n", number)
}