package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Hand struct {
	cards string
	bid   int
	hType HandType
}

type HandType int32

const (
	UNDEFINED HandType = iota
	FIVE_OF_A_KIND
	FOUR_OF_A_KIND
	FULL_HOUSE
	THREE_OF_A_KIND
	TWO_PAIR
	ONE_PAIR
	HIGH_CARD
)

func (h HandType) String() string {
	switch h {
	case UNDEFINED:
		return "UNDEFINED"
	case FIVE_OF_A_KIND:
		return "FIVE_OF_A_KIND"
	case FOUR_OF_A_KIND:
		return "FOUR_OF_A_KIND"
	case FULL_HOUSE:
		return "FULL_HOUSE"
	case THREE_OF_A_KIND:
		return "THREE_OF_A_KIND"
	case TWO_PAIR:
		return "TWO_PAIR"
	case ONE_PAIR:
		return "ONE_PAIR"
	case HIGH_CARD:
		return "HIGH_CARD"
	}
	return ""
}

func findType(m map[string]int) HandType {
	if len(m) == 1 {
		return FIVE_OF_A_KIND
	}
	if len(m) == 2 {
		for _, v := range m {
			if v == 4 {
				return FOUR_OF_A_KIND
			}
			if v == 3 {
				return FULL_HOUSE
			}
		}
	}
	if len(m) == 3 {
		for _, v := range m {
			if v == 3 {
				return THREE_OF_A_KIND
			}
			if v == 2 {
				return TWO_PAIR
			}
		}
	}
	if len(m) == 4 {
		return ONE_PAIR
	}
	if len(m) == 5 {
		return HIGH_CARD
	}
	return UNDEFINED
}

func convert(a byte) byte {
}

func cmpCards(a, b string) int {
	for i := range a {
		if a[i] == b[i] {
			continue
		}

		if unicode.IsLetter(rune(a[i])) && unicode.IsLetter(rune(b[i])) {

		}

		return cmp.Compare(a[i], b[i])

	}
	return 0
}

func CmpHands(a, b Hand) int {
	if b.hType == a.hType {
		return cmpCards(b.cards, a.cards)
	}
	return cmp.Compare(b.hType, a.hType)
}
func main() {

	file, e := os.Open("resources/input2")

	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	answer := 0
	hands := []Hand{}

	// Read file
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		bid, _ := strconv.Atoi(fields[1])
		hand := Hand{fields[0], bid, 0}
		hands = append(hands, hand)
	}

	var m map[string]int

	for i := range hands {
		m = make(map[string]int)
		for _, v := range hands[i].cards {
			m[string(v)] = m[string(v)] + 1
		}
		hands[i].hType = findType(m)
	}

	slices.SortFunc(hands, CmpHands)

	for i := range hands {
		// fmt.Println(hands[i].cards + ": " + hands[i].hType.String() + ", bid: " + strconv.Itoa(hands[i].bid) + ", index: " + strconv.Itoa(i))
		answer += (i + 1) * hands[i].bid
	}
	fmt.Println(answer)

}
