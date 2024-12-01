package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Start")

	file, e := os.Open("resources/input.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	partOne(*scanner)
	partTwo(*scanner)
}

func partTwo(scanner bufio.Scanner) {
	first := []int{}
	second := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Fields(line)
		firstNumber, _ := strconv.Atoi(token[0])
		secondNumber, _ := strconv.Atoi(token[1])
		first = append(first, firstNumber)
		second[secondNumber]++
	}

	sum := 0

	for _, number := range first {
		similiratyScore := number * second[number]
		sum += similiratyScore
	}

	fmt.Println(int64(sum))

}

func partOne(scanner bufio.Scanner) {
	first := []int{}
	second := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Fields(line)
		firstNumber, _ := strconv.Atoi(token[0])
		secondNumber, _ := strconv.Atoi(token[1])
		first = append(first, firstNumber)
		second = append(second, secondNumber)
	}

	slices.Sort(first)
	slices.Sort(second)

	sum := 0.0
	for i := 0; i < len(first); i++ {
		diff := math.Abs(float64(first[i] - second[i]))
		sum += diff

	}

	fmt.Println(int64(sum))

}

