package main

import (
	"bufio"
	"fmt"
	"log"
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
	partTwo(*scanner)
}

type Updates struct {
	element []int
}

func partOne(scanner bufio.Scanner) {
	isSecondPart := false
	pageOrderingRules := make(map[int][]int)
	updatesList := []Updates{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isSecondPart = true
			continue
		}
		if !isSecondPart {
			tmp := strings.Split(line, "|")
			t1, _ := strconv.Atoi(tmp[0])
			t2, _ := strconv.Atoi(tmp[1])
			pageOrderingRules[t1] = append(pageOrderingRules[t1], t2)
		} else {
			tmp := strings.Split(line, ",")
			tmpInt := []int{}
			for _, ele := range tmp {
				t, _ := strconv.Atoi(ele)
				tmpInt = append(tmpInt, t)
			}
			updatesList = append(updatesList, Updates{tmpInt})
		}

	}
	sum := 0
	for _, updates := range updatesList {
		readList := []int{}
		isUpdateCorrect := true
		for _, update := range updates.element {
			check := pageOrderingRules[update]
			for _, toCheck := range check {
				isInReadList := slices.Contains(readList, toCheck)
				if isInReadList {
					isUpdateCorrect = false
					break
				}
			}

			readList = append(readList, update)
		}
		if isUpdateCorrect {
			result := updates.element[len(updates.element)/2]
			sum += result
		}
	}
	fmt.Println(sum)
}

func partTwo(scanner bufio.Scanner) {
	isSecondPart := false
	pageOrderingRules := make(map[int][]int)
	updatesList := []Updates{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isSecondPart = true
			continue
		}
		if !isSecondPart {
			tmp := strings.Split(line, "|")
			t1, _ := strconv.Atoi(tmp[0])
			t2, _ := strconv.Atoi(tmp[1])
			pageOrderingRules[t1] = append(pageOrderingRules[t1], t2)
		} else {
			tmp := strings.Split(line, ",")
			tmpInt := []int{}
			for _, ele := range tmp {
				t, _ := strconv.Atoi(ele)
				tmpInt = append(tmpInt, t)
			}
			updatesList = append(updatesList, Updates{tmpInt})
		}

	}
	sum := 0
	for _, updates := range updatesList {
		isUpdateCorrect := false
		isPerfect := true
		for !isUpdateCorrect {
			isUpdateCorrect = true
			readList := []int{}
			fmt.Println("updates", updates)

			for index, update := range updates.element {
				fmt.Println("#", update)
				check := pageOrderingRules[update]
				for _, toCheck := range check {
					isInReadList := slices.Contains(readList, toCheck)
					if isInReadList {
						isUpdateCorrect = false
						isPerfect = false
						for idy, update := range updates.element {
							if update == toCheck {
								updates.element[index], updates.element[idy] = updates.element[idy], updates.element[index]
								break
							}
						}
						break
					}
				}
				readList = append(readList, update)
			}
		}

		fmt.Println("updates", updates)
		if !isPerfect {
			result := updates.element[len(updates.element)/2]
			sum += result
		}
	}
	fmt.Println(sum)
}
