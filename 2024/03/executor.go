package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	fmt.Println("Start")

	file, e := os.Open("resources/input.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	partTwo(*scanner)
}

func partOne(scanner bufio.Scanner) {
	content := ""
	for scanner.Scan() {
		line := scanner.Text()
		content += line
	}
	r, _ := regexp.Compile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
	c := r.FindAllStringSubmatch(content, -1)
	sum := 0
	for _, matches := range c {
		// name := matches[0]
		firstNumber, _ := strconv.Atoi(matches[1])
		secondNumber, _ := strconv.Atoi(matches[2])
		resultOfMatch := firstNumber * secondNumber
		sum += resultOfMatch
	}
	fmt.Println(sum)
}

func partTwo(scanner bufio.Scanner) {

	content := ""
	for scanner.Scan() {
		line := scanner.Text()
		content += line
	}
	r, _ := regexp.Compile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)|do\\(\\)|don\\'t\\(\\)")
	c := r.FindAllStringSubmatch(content, -1)
	sum := 0
	do := true
	for _, matches := range c {
		name := matches[0]
		if name == "do()" {
			do = true
		}
		if name == "don't()" {
			do = false
		}
		firstNumber, _ := strconv.Atoi(matches[1])
		secondNumber, _ := strconv.Atoi(matches[2])
		if do {
			resultOfMatch := firstNumber * secondNumber
			sum += resultOfMatch
		}
	}
	fmt.Println(sum)
}
