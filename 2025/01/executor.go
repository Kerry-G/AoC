package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Direction int

const (
	Left Direction = iota
	Right
)

var directionName = map[Direction]string{
	Left:  "Left",
	Right: "Right",
}

func (ss Direction) String() string {
	return directionName[ss]
}

func parseDirection(s string) Direction {
	if s == "L" {
		return Left
	} else {
		return Right
	}
}

type Step struct {
	direction Direction
	distance  int
}

type Dial struct {
	current int
}

func (dial *Dial) turn(step Step) {
	if step.direction == Left {
		dial.current -= step.distance
	} else {
		dial.current += step.distance
	}
	dial.current = dial.current % 100
}

func main() {

	fmt.Println("Start")

	file, e := os.Open("resources/input.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	partOne(*scanner)
	// partTwo(*scanner)
}

func partOne(scanner bufio.Scanner) {

	steps := []Step{}

	for scanner.Scan() {
		line := scanner.Text()
		distance, _ := strconv.Atoi(line[1:])
		step := Step{parseDirection(line[0:1]), distance}
		steps = append(steps, step)
	}

	dial := Dial{50}
	sum := 0
	for _, step := range steps {
		dial.turn(step)
		if dial.current == 0 {
			sum += 1
		}
	}
	println(sum)

}

func partTwo(scanner bufio.Scanner) {
	for scanner.Scan() {

	}
}
