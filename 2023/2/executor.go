package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	const RED int = 12
	const GREEN int = 13
	const BLUE int = 14

	answer := 0

	file, err := os.Open("resources/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameIdRegex := regexp.MustCompile("[0-9]+")
		gameId, _ := strconv.Atoi(gameIdRegex.FindString(line))
		sets := strings.Split(line, ";")
		gameGood := true
		for _, set := range sets {
			regex := regexp.MustCompile("[0-9]+ (green|blue|red)")
			a := regex.FindAllString(set, -1)

			blue, green, red := 0, 0, 0
			for _, val := range a {
				entry := strings.Split(val, " ")
				count, _ := strconv.Atoi(entry[0])
				color := entry[1]
				switch color {
				case "blue":
					blue += count
				case "red":
					red += count
				case "green":
					green += count
				}

			}

			if blue > BLUE || red > RED || green > GREEN {
				gameGood = false
			}
		}
		if gameGood {
			answer += gameId
		}
	}

	fmt.Println(answer)
}
