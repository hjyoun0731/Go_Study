package main

import "fmt"

func main() {

	var clothes [][]string

	clothes = [][]string{{"yellow_hat", "headgear"}, {"blue_sunglasses", "eyewear"}, {"green_turban", "headgear"}}

	ret := solution(clothes)
	fmt.Println(ret)
}

func solution(clothes [][]string) int {

	var c map[string]int
	var ret int

	ret = 1
	c = make(map[string]int)

	for index := range clothes {
		c[clothes[index][1]] = 1
	}
	for index := range clothes {
		c[clothes[index][1]]++
	}
	for kinds := range c {
		ret *= c[kinds]
	}
	ret--

	return ret
}
