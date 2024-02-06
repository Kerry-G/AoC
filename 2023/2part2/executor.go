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

	answer := 0

	file, err := os.Open("resources/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sets := strings.Split(line, ";")
		minB, minR, minG := -1, -1, -1
		for _, set := range sets {
			regex := regexp.MustCompile("[0-9]+ (green|blue|red)")
			a := regex.FindAllString(set, -1)

			blue, green, red := -1, -1, -1
			for _, val := range a {
				entry := strings.Split(val, " ")
				count, _ := strconv.Atoi(entry[0])
				color := entry[1]
				switch color {
				case "blue":
					blue = count
				case "red":
					red = count
				case "green":
					green = count
				}

				if (minB < blue || minB == -1) && blue != -1 {
					minB = blue
				}
				if (minG < green || minG == -1) && green != -1 {
					minG = green
				}
				if (minR < red || minR == -1) && red != -1 {
					minR = red
				}

			}
		}

		power := minR * minG * minB
		answer += power
	}

	fmt.Println(answer)
}
