package main

import "fmt"

func main() {
	var participant []string
	var completion []string

	participant = []string{"mislav", "stanko", "mislav", "ana"}
	completion = []string{"stanko", "ana", "mislav"}

	answer := solution(participant, completion)
	fmt.Println(answer)
}

func solution(participant []string, completion []string) string {
	var answer string

	var p map[string]int

	p = make(map[string]int)

	for i := 0; i < len(participant); i++ {
		if participant[i] != "\n" {
			p[participant[i]]++
		} else {
			break
		}
	}
	for i := 0; i < len(completion); i++ {
		if completion[i] != "\n" {
			p[completion[i]]--
		} else {
			break
		}
	}
	for i := 0; i < len(participant); i++ {
		if p[participant[i]] == 1 {
			answer = participant[i]
			break
		} else {
		}
	}

	return answer
}
