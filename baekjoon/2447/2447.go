package main

import (
	"bufio"
	"fmt"
	"os"
)

func makeStar(n int) [][]string {
	stars := make([][]string, n)
	for i := 0; i < n; i++ {
		stars[i] = make([]string, n)
	}
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if x%3 == 1 && y%3 == 1 {
				stars[x][y] = " "
			} else {
				stars[x][y] = "*"
			}
		}
	}

	return stars
}

func makeBlank(n int, blankSize int, p [][]string) [][]string {
	if blankSize == 1 {
		return p
	}
	for i := blankSize; i < n; i += blankSize * 3 {
		for j := blankSize; j < n; j += blankSize * 3 {
			for x := i; x < i+blankSize; x++ {
				for y := j; y < j+blankSize; y++ {
					p[x][y] = " "
				}
			}
		}
	}
	return makeBlank(n, blankSize/3, p)
}

func printStar(n int) {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	pattern := makeStar(n)
	pattern = makeBlank(n, n/3, pattern)

	for i := 0; i < n; i++ {
		for _, i := range pattern[i] {
			writer.WriteString(i)
		}
		writer.WriteString("\n")
	}
}

func main() {
	var N int
	_, _ = fmt.Scan(&N)
	printStar(N)
}