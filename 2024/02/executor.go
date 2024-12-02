package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	level                Level
	possibleLevel        []Level
	isValid              bool
	isPossibleLevelValid bool
}

type Level struct {
	elements []int
}

func (report *Report) addElement(element int) {
	report.level.addElement(element)
}

func (report *Report) addAllPossibleLevel() {
	for index, _ := range report.level.elements {
		tempLevel := Level{}
		if index == 0 {
			tempLevel = Level{report.level.elements[1:]}
		} else if index == len(report.level.elements)-1 {
			tempLevel = Level{report.level.elements[:len(report.level.elements)-1]}
		} else {
			tempLevel = Level{append([]int{}, report.level.elements[:index]...)}
			tempLevel.elements = append(tempLevel.elements, report.level.elements[index+1:]...)
		}
		report.addPossibleLevel(tempLevel)
	}
}

func (report *Report) addPossibleLevel(level Level) {
	report.possibleLevel = append(report.possibleLevel, level)
}

func (level *Level) isReportValid() (success bool, index int) {
	lastLevelSeen := -1
	reportType := "UNDEFINED"
	for index, level := range level.elements {
		if lastLevelSeen == -1 {
			lastLevelSeen = level
			continue
		}

		difference := AbsDifference(lastLevelSeen, level)
		if difference < 1 || difference > 3 {
			return false, index
		}

		if lastLevelSeen > level {
			if reportType == "UNDEFINED" {
				reportType = "DECREASING"
			}
			if reportType == "INCREASING" {
				return false, index
			}
		} else {
			if reportType == "UNDEFINED" {
				reportType = "INCREASING"
			}
			if reportType == "DECREASING" {
				return false, index
			}
		}

		lastLevelSeen = level
	}
	return true, -1
}
func (report *Report) isReportValidWithDamper() bool {
	for _, level := range report.possibleLevel {
		isValid, _ := level.isReportValid()
		if isValid {
			report.isPossibleLevelValid = true
			return true
		}
	}
	report.isPossibleLevelValid = false
	return false
}

func (level *Level) addElement(element int) {
	level.elements = append(level.elements, element)
}

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

	reports := []Report{}
	for scanner.Scan() {
		report := Report{}
		line := scanner.Text()
		tokens := strings.Fields(line)
		for _, token := range tokens {
			intToken, _ := strconv.Atoi(token)
			report.addElement(intToken)
		}
		reports = append(reports, report)
	}

	sum := 0
	for _, report := range reports {
		isValid, _ := report.level.isReportValid()

		if isValid {
			sum++
		}

	}

	fmt.Println(sum)
}

func partTwo(scanner bufio.Scanner) {

	reports := []Report{}
	for scanner.Scan() {
		report := Report{}
		line := scanner.Text()
		tokens := strings.Fields(line)
		for _, token := range tokens {
			intToken, _ := strconv.Atoi(token)
			report.addElement(intToken)
		}
		reports = append(reports, report)
	}

	sum := 0
	for _, report := range reports {

		isValid, _ := report.level.isReportValid()

		if isValid {
			sum++
		} else {
			report.addAllPossibleLevel()
			if report.isReportValidWithDamper() {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func AbsDifference(a, b int) int {
	return Abs(a - b)
}
