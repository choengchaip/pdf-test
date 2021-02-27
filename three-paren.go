package main

import (
	"log"
)

func main() {
	problems := []string{"(())", "(()))(()", "(()())()", "({])[]", "[({}])", "()[{}()]"}
	for _, problem := range problems {
		log.Println(solve(problem))
	}
}

var (
	charMatches = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
)

func solve(problem string) bool {
	stack := make([]string, 0)
	for i := range problem {
		symbol := string(problem[i])
		if len(stack) == 0 {
			stack = append(stack, symbol)
		} else {
			if stack[len(stack)-1] == charMatches[symbol] {
				stack = popOutArray(stack)
			} else {
				stack = append(stack, symbol)
			}
		}
	}
	return len(stack) == 0
}

func popOutArray(s []string) []string {
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}
