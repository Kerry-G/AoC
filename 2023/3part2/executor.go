package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"unicode"
)

type Item struct {
	Number    string
	start_pos [2]int
	len       int
}

type Gear struct {
	position [2]int
	r1       Item
	r2       Item
	ratio    int
}

func getItemAtCoord(coord [2]int, world [][]rune) Item {
	x, y := coord[0], coord[1]
	limit_x := len(world[0]) - 1

	left, right := x, x
	for {
		left--
		if left < 0 {
			left = 0
			break
		}

		if !unicode.IsDigit(world[y][left]) {
			left++
			break
		}
	}
	for {
		right++
		if right > limit_x {
			right = limit_x
			break
		}
		if !unicode.IsDigit(world[y][right]) {
			right--
			break
		}
	}

	ans := ""
	for k := left; k <= right; k++ {
		ans += string(world[y][k])
	}

	item := Item{ans, [2]int{left, y}, right - left}
	return item
}

func scanAdjacent(i *Gear, world [][]rune) {

	sr, sc, er, ec := i.position[0], i.position[1], i.position[0]+1, i.position[1]+1
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

	items := []Item{}
	for k := sr; k <= er; k++ {
		for l := sc; l <= ec; l++ {
			char := world[l][k]
			if unicode.IsDigit(char) {
				item := getItemAtCoord([2]int{k, l}, world)
				itemsContainsItem := slices.ContainsFunc(items, func(i Item) bool {
					return i.start_pos[0] == item.start_pos[0] && i.start_pos[1] == item.start_pos[1]
				})
				if itemsContainsItem {
					continue
				}
				items = append(items, item)
			}
		}
	}

	if len(items) == 2 {
		r1, _ := strconv.Atoi(items[0].Number)
		r2, _ := strconv.Atoi(items[1].Number)
		i.r1 = items[0]
		i.r2 = items[1]
		i.ratio = r1 * r2
	}
}

func main() {

	file, e := os.Open("resources/input")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	items := []*Item{}
	gears := []*Gear{}
	itr := 0
	column := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		newItem := false
		var i *Item
		for idx, val := range line {
			row = append(row, val)
			if string(val) == "*" {
				newGear := &Gear{position: [2]int{idx, itr}}
				gears = append(gears, newGear)
				continue
			}
			if unicode.IsDigit(val) {
				if !newItem {
					newItem = true
					i = &Item{string(val), [2]int{idx, itr}, 1}
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
	for _, g := range gears {
		scanAdjacent(g, column)
		// fmt.Println("Gear position: ", *&g.position, "| r1: ", g.r1.Number, " | r2: ", g.r2.Number, "| ratio: ", g.ratio)
		sum += g.ratio
	}
	fmt.Println("Sum: ", sum)
}
