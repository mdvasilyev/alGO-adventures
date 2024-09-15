package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Team struct {
	name          string
	failures      map[rune]int
	successNumber int
	penalty       int
}

func getTime(start string, cur string) int {
	hackathonStart, _ := time.Parse("15:04:05", start)
	current, _ := time.Parse("15:04:05", cur)
	if current.Before(hackathonStart) {
		current = current.Add(24 * time.Hour)
	}
	return int(current.Sub(hackathonStart).Minutes())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	start := scanner.Text()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	teams := make(map[string]*Team)

	for i := 0; i < n; i++ {
		scanner.Scan()
		str := scanner.Text()
		strs := strings.Split(str, " ")
		name := strings.Trim(strs[0], "\"")
		requestTime := strs[1]
		server := rune(strs[2][0])
		action := strs[3]

		if _, ok := teams[name]; !ok {
			teams[name] = &Team{
				name:     name,
				failures: make(map[rune]int),
			}
		}
		team := teams[name]

		switch action {
		case "ACCESSED":
			minutes := getTime(start, requestTime)
			team.successNumber++
			team.penalty += minutes + team.failures[server]*20
		case "DENIED", "FORBIDEN":
			team.failures[server]++
		case "PONG":
		}
	}

	var results []*Team
	for _, team := range teams {
		results = append(results, team)
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].successNumber != results[j].successNumber {
			return results[i].successNumber > results[j].successNumber
		}
		if results[i].penalty != results[j].penalty {
			return results[i].penalty < results[j].penalty
		}
		return results[i].name < results[j].name
	})

	for i, team := range results {
		place := i + 1
		if i > 0 && results[i].successNumber == results[i-1].successNumber && results[i].penalty == results[i-1].penalty {
			place = i
		}
		fmt.Printf("%d \"%s\" %d %d\n", place, team.name, team.successNumber, team.penalty)
	}
}
