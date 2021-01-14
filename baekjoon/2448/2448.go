package main

import (
	"fmt"
	"math"
)

var pattern []string

func makeBlank(k int, n int, star []string) []string {
	for i :=0;i+3<n;i+=3 {
		for j:=0;j<n-(i+3);j++ {
			star[i] += " "
			star[i+1] += " "
			star[i+2] += " "
		}
	}

	return star
}

func makeStar(k int, n int, star []string) []string {

	if int(math.Pow(2, float64(k))) > n/3 {
		return star
	}

	if k==0 {
		pattern[0]="  *   "
		pattern[1]=" * *  "
		pattern[2]="***** "
		star[0]+=pattern[0]
		star[1]+=pattern[1]
		star[2]+=pattern[2]
		return makeStar(k+1, n, star)
	}
	for count := 0; count <= k; count++ {

		for i := 0; i < 3; i++ {
			switch i {
			case 0:
				star[k*3+i] += pattern[0]
			case 1:
				star[k*3+i] += pattern[1]
			case 2:
				star[k*3+i] += pattern[2]
			}
		}
	}


	return makeStar(k+1, n, star)
}

func printStar(k int, n int) {
	var stars string
	var star []string
	pattern = make([]string,n)
	star = make([]string,n)
	star = makeBlank(k,n,star)
	star = makeStar(k, n, star)
	for _,line := range star{
		stars+=line
		stars+="\n"
	}
	fmt.Print(stars)
}

func main() {
	N := []int{3, 6, 12, 24, 48, 96}
	var k int
	k = 2
	printStar(0, N[k])
}
