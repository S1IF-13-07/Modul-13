package main

import "fmt"

func main() {
	var username, password string
	failCount := 0
	correctUser := "Admin"
	correctPass := "Admin"
	isCorrect := false

	for !isCorrect {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)

		if username == correctUser && password == correctPass {
			isCorrect = true
		} else {
			failCount++
		}
	}

	fmt.Printf("%d percobaan gagal login\n", failCount)
}
