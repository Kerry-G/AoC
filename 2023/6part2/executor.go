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

	game := Game{}
	time := ""
	distance := ""
	// Read file
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		mode := -1
		for _, f := range fields {
			if strings.Contains(f, "Time") {
				mode = 0
				continue
			}
			if strings.Contains(f, "Distance") {
				mode = 1
				continue
			}


			switch mode {
			case 0:
				time = time + f
			case 1:
				distance = distance + f
			}

		}
	}
	game.time, _ = strconv.Atoi(time)
	game.distance, _ = strconv.Atoi(distance)
	ctr := 0
	// brute force this. 
	// computer is fast enough to deal with this
	// a better way would be to find the limit of acceptable with a binary search
	for i := range game.time {
		if verify(i, game.time, game.distance) {ctr++}
	}
	fmt.Println(ctr)
}
