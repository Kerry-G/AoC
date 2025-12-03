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

func (ss Step) String() string {
	return ss.direction.String() + " " + strconv.Itoa(ss.distance)
}

type Dial struct {
	current int
}

func (dial *Dial) turn(step Step) int {
	// Handle step distance over 100 to a 2 digit distance and keep the amount of click it would add
	click := step.distance / 100
	if click < 0 {
		click = -click
	}
	step.distance = step.distance % 100

	temp := dial.current
	if step.direction == Left {
		dial.current -= step.distance
	} else {
		dial.current += step.distance
	}

	if dial.current/100 == 1 && dial.current != 100 {
		click += 1
	}
	dial.current = dial.current % 100

	// if dial.current went negative, set it back to a normal number
	if dial.current < 0 {
		dial.current = 100 + dial.current
		if temp != 0 {
			click += 1
		}
	}

	// add a click if it did finish with 0
	if dial.current == 0 {
		click += 1
	}

	return click
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
		fmt.Println("Step: ", step)
		click := dial.turn(step)
		fmt.Println("Dial counter:", dial.current)
		sum += click
		fmt.Println("Click: ", click)
	}
	println(sum)

}
