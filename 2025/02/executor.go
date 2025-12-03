package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	start int
	end   int
}

func getLength(s int) int {
	return len(strconv.Itoa(s))
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
	entries := []Entry{}

	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Split(line, ",")
		for _, entry := range token {
			temp := strings.Split(entry, "-")
			start, _ := strconv.Atoi(temp[0])
			end, _ := strconv.Atoi(temp[1])
			entries = append(entries, Entry{start, end})
		}
	}
	sum := 0
	for _, entry := range entries {
		fmt.Println(entry)
		for i := entry.start; i <= entry.end; i++ {
			len := getLength(i)
			if len%2 != 0 {
				continue
			}
			slicer := int(math.Pow(10, float64(len/2)))
			firstHalf := (i / slicer)
			secondHalf := (i % slicer)
			if firstHalf == secondHalf {
				sum += i
			}
		}
	}
	fmt.Println(sum)

}
