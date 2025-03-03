package main

import "fmt"

func generate(cur string, opened, closed, total int) []string {
	strings := make([]string, 0)

	if len(cur) == 2*total {
		return []string{cur}
	}

	if opened < total {
		strings = append(strings, generate(cur+"(", opened+1, closed, total)...)
	}

	if closed < opened {
		strings = append(strings, generate(cur+")", opened, closed+1, total)...)
	}

	return strings
}

func generateParenthesis(n int) []string {
	return generate("", 0, 0, n)
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
}
