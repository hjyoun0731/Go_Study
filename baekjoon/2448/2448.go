package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var pattern []string

func makeBlank(kkm int, star []string) []string {
	for i := 0; i < kkm; i++ {
		for j := 0; j < kkm; j = j + 3 {
			star[i] = "   " + star[i] + "   "
		}
	}

	return star
}

func makeStar(k int, l int, n int, star []string, pattern []string) []string {
	if (l+1)*3 > n {
		return star
	}

	if k == 0 {
		star[0] = pattern[0]
		star[1] = pattern[1]
		star[2] = pattern[2]
		return makeStar(k+1, 1, n, star, pattern)
	}

	kk := int(math.Pow(2, float64(k))) * 3
	kkm := int(math.Pow(2, float64(k-1))) * 3

	star = makeBlank(kkm, star)

	j := 0
	for i := kkm; i < kk; i++ {
		star[i] = pattern[j] + pattern[j]
		j++
	}
	copy(pattern, star)
	return makeStar(k+1, l*2+1, n, star, pattern)
}

func printStar(k int, l int, n int) {
	var star []string
	pattern = make([]string, n)
	star = make([]string, n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	pattern[0] = "  *   "
	pattern[1] = " * *  "
	pattern[2] = "***** "

	star = makeStar(k, l, n, star, pattern)
	stars := strings.Join(star, "\n")
	writer.WriteString(stars)
}

func main() {
	var n int
	fmt.Scan(&n)
	printStar(0, 0, n)
}
