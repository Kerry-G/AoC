package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Case struct {
	id             int
	ownNumbers     map[int]bool
	winningNumbers map[int]bool
}

func readCaseFromString(s string) Case {
	token := strings.Split(s, " ")
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

		if len(t) == 0 {
			// Ignore empty spaces that happen when there is two spaces in the original string
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
	c := Case{id, ownNumbers, winningNumbers}
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

	points := 0.0
	for _, c := range cases {
		matchNumbers := intersect(c.winningNumbers, c.ownNumbers)
		if len(matchNumbers) == 0 {
			// no match, better luck next time
			continue
		}
		points += math.Pow(2, float64(len(matchNumbers)-1))
	}
	fmt.Println("Points: ", points)
}
