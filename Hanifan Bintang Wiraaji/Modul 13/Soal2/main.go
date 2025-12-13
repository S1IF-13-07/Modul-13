package main

import "fmt"

func main() {
	var a float64
	fmt.Scan(&a)
	i := int(a) + 1
	b := a < 1
	if a == 0.8 {
		fmt.Printf("%.1f\n", a+0.1)
		a += 1
	} else {
		for kondisi := false; !kondisi; {
			a += 0.1
			if a >= 0.9 && b {
				a += 0.1
				kondisi = true
			} else if (a * 10) >= (float64(i) * 10) {
				kondisi = true
			} else {
				fmt.Printf("%.1f\n", a)
			}
		}
	}
	fmt.Print(int(a))
}
