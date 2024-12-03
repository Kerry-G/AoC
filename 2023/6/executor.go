package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	time     int
	distance int
	done     int
}

func verify(hold int, time int, distance int) bool {
	travel_time := time-hold
	distance_travelled := travel_time * hold
	if distance_travelled > distance { return true }
	return false
}

func main() {

	file, e := os.Open("resources/input2")

	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	games := []Game{}
	
	// Read file
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		idx := 0
		mode := -1
		for _, f := range fields {
			if strings.Contains(f, "Time") {
				idx = 0
				mode = 0
				continue
			}
			if strings.Contains(f, "Distance") {
				idx = 0
				mode = 1
				continue
			}

			value, _ := strconv.Atoi(f)

			switch mode {
			case 0:
				game := Game{time: value}
				games = append(games, game)
			case 1:
				games[idx].distance = value
			}
			idx = idx + 1

		}
	}

	answer := 1
	for i := range games {
		ctr := 0 
		for v := range games[i].time {
			if verify(v, games[i].time, games[i].distance) { ctr++ }
		}
		games[i].done = ctr
		answer = answer * ctr
	}

	fmt.Println(answer)

}
