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

func reverse(s string) string {
	rline := ""
	for i := len(s) - 1; i >= 0; i-- {
		rline += string(s[i])
	}
	return rline

}

func main() {

	acc := 0
	file, err := os.Open("resources/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		regex := regexp.MustCompile(`(([1-9])|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine))`)
		digit := regex.FindString(line)

		rregex := regexp.MustCompile(`(([1-9])|(eno)|(owt)|(eerht)|(ruof)|(evif)|(xis)|(neves)|(thgie)|(enin))`)
		rdigit := reverse(rregex.FindString(reverse(line)))

		r := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
		firstDigit := r.Replace(digit)
		lastDigit := r.Replace(rdigit)
		ans := firstDigit + lastDigit
		tempacc, _ := strconv.Atoi(ans)
		acc += tempacc
	}

	fmt.Println(acc)

}
