package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Item struct {
	Number    string
	Adjacent  bool
	start_pos [2]int
	len       int
}

func scanAdjacent(i *Item, world [][]rune) {

	// If we have start_pos 3,3 and len 3, we wnat to verify the scope of 2,2 => 2,7; 3,2 => 3,7 and 4,2 => 4,7.
	sr, sc, er, ec := i.start_pos[0], i.start_pos[1], i.start_pos[0]+i.len+1, i.start_pos[1]+1
	limitr, limitc := len(world[0])-1, len(world)-1
	if sr != 0 {
		sr = sr - 1
	}
	if sc != 0 {
		sc = sc - 1
	}
	if er > limitr {
		er = limitr
	}
	if ec > limitc {
		ec = limitc
	}
	// fmt.Print("\n", i.Number, ": ")
	// fmt.Print(sr, " to: ", er, ", ")
	// fmt.Println(sc, " to: ", ec)
	var dot rune = 46
	for k := sr; k <= er; k++ {
		for l := sc; l <= ec; l++ {
			char := world[l][k]
			if unicode.IsDigit(char) {
				continue
			}

			if char != dot {
				i.Adjacent = true
				return
			}
		}
	}

}

func main() {

	fmt.Println("Start")

	file, e := os.Open("resources/input10")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	items := []*Item{}
	itr := 0
	column := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		newItem := false
		var i *Item
		for idx, val := range line {
			row = append(row, val)
			if unicode.IsDigit(val) {
				if !newItem {
					newItem = true
					i = &Item{string(val), false, [2]int{idx, itr}, 1}
					if idx == len(line)-1 {
						newItem = false
						items = append(items, i)
					}
				} else {
					i.len++
					i.Number += string(val)
					if idx == len(line)-1 {
						newItem = false
						items = append(items, i)
					}
				}
			} else {
				if newItem {
					newItem = false
					items = append(items, i)
				}
			}

		}
		column = append(column, row)
		itr++
	}

	sum := 0
	for _, it := range items {
		scanAdjacent(it, column)
		// fmt.Println("\t\t", it.Number, it.Adjacent)
		if it.Adjacent {
			nb, _ := strconv.Atoi(it.Number)
			sum += nb
		}
	}
	fmt.Println("Sum: ", sum)
}
