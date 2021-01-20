package main

import (
	"fmt"

	"github.com/Go_Study/nomadcoder/learnDict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	err := dictionary.Add(baseWord, "First")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(baseWord)

	err = dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	dictionary.Delete(baseWord)
	word, _ = dictionary.Search(baseWord)
	fmt.Println(word)

}
