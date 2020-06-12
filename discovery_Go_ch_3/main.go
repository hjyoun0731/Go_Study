package main

import (
	"fmt"
	"hangul"
)

func chapter3() {
	for i, r := range "가나다" {
		fmt.Println(i, r)
	}
	fmt.Println(len("가나다"))
}

func main() {
	chapter3()
	hangul.HasConsonantSuffix("테스트")
}
