package main  
import "fmt"  
func main() { 

	var kata string 
	var pengulangan int 
	fmt.Scan(&kata, &pengulangan) 
	kondisi := 0 
	for kondisi1 := false; !kondisi1; { 
		fmt.Println(kata) 
		kondisi++ 
		kondisi1 = (kondisi >= pengulangan) 
	} 
} 