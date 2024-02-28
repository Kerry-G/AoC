package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Mapper struct {
	sourceName      string
	destinationName string
	ranges          [][3]int
}

func parse(file *os.File) (seeds []int, mappers []Mapper) {

	scanner := bufio.NewScanner(file)

	lineCounter := 0
	mapper := Mapper{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// At lineCounter 0, we read the seeds
		if lineCounter == 0 {
			for i, f := range fields {
				if i == 0 {
					continue
				}
				seed, _ := strconv.Atoi(f)
				seeds = append(seeds, seed)
			}
		} else if lineCounter == 1 {
			lineCounter++
			continue
		} else if len(fields) == 0 {
			mappers = append(mappers, mapper)
			mapper = Mapper{}
		} else if len(fields) == 2 {
			token := strings.Split(fields[0], "-")
			mapper.sourceName = token[0]
			mapper.destinationName = token[2]
		} else if len(fields) == 3 {
			olivier := [3]int{}
			for i, element := range fields {
				olivier[i], _ = strconv.Atoi(element)
			}
			mapper.ranges = append(mapper.ranges, olivier)
		}

		lineCounter++
	}
	// There is no white space at the last line, thus we need to save the last mapper that we read
	mappers = append(mappers, mapper)
	return seeds, mappers
}

func main() {
	file, e := os.Open("resources/input2")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()
	seeds, mappers := parse(file)
	answer := -1
	for _, seed := range seeds {
		fmt.Println("Seeds: ", seed)
		result := seed
		for _, mapper := range mappers {
			fmt.Print(mapper.sourceName, "=>", mapper.destinationName, ": ")
			for _, olivier := range mapper.ranges {

				OG := olivier[1]
				FG := OG + olivier[2] - 1
				offset := olivier[0] - OG
				if result >= OG && result <= FG {
					result = result + offset
					break
				}
			}
			fmt.Println(result)
		}
		fmt.Println("=======================================")

		if result < answer || answer == -1 {
			answer = result
		}
	}

	fmt.Println("The answer is: ", answer)

}
