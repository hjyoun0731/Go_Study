package main

import (
	"fmt"
	"math"
)

func main() {
	var n []int
	for i := 1; i < 8; i++ {
		n = append(n, int(math.Pow(3, float64(i))))
	}
	fmt.Println(n)

	const k = 1
	var stars []string

	for y := 1; y <= n[k]; y++ {
		var star string
		switch {
		case y%9%3 == 2:
			for x := 1; x <= n[k]; x += 3 {
				star += "* *"
			}
		default:
			for x := 1; x <= n[k]; x += 3 {
				star += "***"
			}
		}
		stars = append(stars, star)
		stars = append(stars, "\n")
	}
	fmt.Print(stars)
}
