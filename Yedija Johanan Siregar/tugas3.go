package main

import "fmt"

func main() {
  var target, donasi, total int
  var jumlahDonatur int = 0
  
  fmt.Scan(&target)

  for {
      fmt.Scan(&donasi)
      jumlahDonatur++
      total += donasi

      fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, total)

    if total >= target {
        break
    }
  }

  fmt.Printf("Target tercapai! Total donasi: %d dari %ddonatur.\n", total, jumlahDonatur)
}
