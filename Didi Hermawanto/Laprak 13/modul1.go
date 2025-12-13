package main
import "fmt"

func main() {
    var word string
    var repetitions int

    fmt.Print("Masukkan kata dan jumlah pengulangan: ")
    fmt.Scan(&word, &repetitions)

    counter := 0
    done := false

    for !done {
        fmt.Println(word)
        counter++
        done = (counter >= repetitions)
    }
}
