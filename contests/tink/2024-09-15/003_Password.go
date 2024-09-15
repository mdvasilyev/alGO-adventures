package main

import (
	"fmt"
	"strings"
)

func validateMap(stMap *map[string]int, m *map[string]int) bool {
	for k := range *m {
		if _, ok := (*stMap)[k]; !ok || (*m)[k] == 0 {
			return false
		}
	}
	return true
}

func main() {
	var history, req string
	var maxLen int

	_, _ = fmt.Scan(&history, &req)
	_, _ = fmt.Scan(&maxLen)

	minLen := len(req)

	if minLen > len(history) {
		fmt.Println("-1")
		return
	}

	reqList := strings.Split(req, "")
	stMap := make(map[string]int)

	for _, v := range reqList {
		stMap[v] = 0
	}

	l, r := len(history)-minLen, len(history)-1
	m := make(map[string]int)

	for _, val := range req {
		m[string(val)] = 0
	}

	for i := l; i <= r; i++ {
		m[string(history[i])]++
	}

	for l+r > 0 {
		if !validateMap(&stMap, &m) {
			if r-l+1 != maxLen {
				m[string(history[l-1])]++
				l--
			} else {
				m[string(history[r])]--
				r--
				for r-l+1 != minLen {
					if r-l+1 > minLen {
						m[string(history[l])]--
						l++
					} else {
						m[string(history[l-1])]++
						l--
					}
				}
			}
		} else {
			for i := l; i <= r; i++ {
				fmt.Print(string(history[i]))
			}
			fmt.Println()
			return
		}
	}

	fmt.Println("-1")
}
