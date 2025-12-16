package main 
import "fmt" 
func main() { 

     var angka int 
     var kondisi bool 
     for kondisi = true; kondisi; { 
         fmt.Scan(&angka) 
         kondisi = angka <= 0 
     } 
     fmt.Printf("%d adalah bilangan bulat positif\n", angka) 
}