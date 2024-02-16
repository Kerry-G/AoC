package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Case struct {
	id             int
	ownNumbers     map[int]bool
	winningNumbers map[int]bool
	amount         int
}

func readCaseFromString(s string) Case {
	// Note to myself: Fields is much more appropriate than Split(s, " ") in this case!
	token := strings.Fields(s)
	id := -1
	ownNumbers := map[int]bool{}
	winningNumbers := map[int]bool{}
	separator := false
	for idx, t := range token {
		if idx == 0 {
			// Ignore first token
			continue
		}
		if idx == 1 {
			// Second token is the id in the form of "1:"
			id, _ = strconv.Atoi(strings.Trim(t, ":"))
			continue
		}
		// back to the main program
		if separator {
			number, _ := strconv.Atoi(t)
			winningNumbers[number] = true
		} else {
			if t == "|" {
				separator = true
				continue
			}
			number, _ := strconv.Atoi(t)
			ownNumbers[number] = true
		}
	}
	c := Case{id, ownNumbers, winningNumbers, 1}
	return c
}

func intersect(m1, m2 map[int]bool) map[int]bool {
	result := map[int]bool{}

	for a, _ := range m1 {
		if m2[a] {
			result[a] = true
		}
	}

	return result
}

func main() {
	file, e := os.Open("resources/input2")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cases := []Case{}
	for scanner.Scan() {
		line := scanner.Text()
		cases = append(cases, readCaseFromString(line))
	}

	nbOfCases := len(cases)
	for idx, _ := range cases {
		matchNumbers := intersect(cases[idx].winningNumbers, cases[idx].ownNumbers)
		nbOfMatches := len(matchNumbers)
		for a := 0; a < cases[idx].amount; a++ {
			for i := 1; i <= nbOfMatches; i++ {
				if (idx + i) > nbOfCases-1 {
					continue
				}
				cases[idx+i].amount++
			}
		}
	}
	answer := 0
	for _, c := range cases {
		fmt.Println("Card id: ", c.id, " amount:", c.amount)
		answer += c.amount
	}
	fmt.Println(answer)
}
