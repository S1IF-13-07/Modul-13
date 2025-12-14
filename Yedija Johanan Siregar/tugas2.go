package main

import "fmt"

func main() {
  var n float64
  fmt.Scan(&n)

  atas := int(n) + 1
    for {
      n += 0.1
    fmt.Printf("%.1f\n", n)

      if int(n) == atas {

      break
    }
  }
}
