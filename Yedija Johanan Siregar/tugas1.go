package main

import "fmt"

func main() {
  var bilangan int
  fmt.Scan(&bilangan)

  count := 0
    
  for {
    count++
  bilangan = bilangan / 10
    if bilangan == 0 {
  
      break
    }
  }
  fmt.Println(count)
}
