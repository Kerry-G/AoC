package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findFirstDigit(s string) string {
	i := strings.IndexAny(s, "123456789")
	return string(s[i])
}

func findLastDigit(s string) string {
	i := strings.LastIndexAny(s, "123456789")
	return string(s[i])
}

func main() {

	file, err := os.Open("resources/input")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	acc := 0
	for scanner.Scan() {
		line := scanner.Text()
		digit := findFirstDigit(line) + findLastDigit(line)
		tempacc, _ := strconv.Atoi(digit)
		acc += tempacc
	}

	fmt.Println(acc)

}
