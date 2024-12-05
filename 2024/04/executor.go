package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {

	fmt.Println("Start")

	file, e := os.Open("resources/test.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	partOne(*scanner)
}

type Node struct {
	element rune
	TL      *Node
	TM      *Node
	TR      *Node
	ML      *Node
	MR      *Node
	BL      *Node
	BM      *Node
	BR      *Node
}

func (node Node) howManyXmas() int {
	fmt.Printf("%c", node.element)
	fmt.Println(node.chain4("MR", ""))
	return 0
}

func (node Node) chain4(direction string, acc string) string {

	v := reflect.ValueOf(node)
	temp := v.FieldByName(direction)
	ele := temp.Interface().(*Node)

	if ele == nil {
		return "."
	}
	if len(acc) == 0 {
		acc = string(ele.element)
	}
	if len(acc) == 4 {
		return acc
	}
	fmt.Println(acc)
	return acc + node.chain4(direction, acc)
}

func partOne(scanner bufio.Scanner) {
	nodes := make([][]Node, 0)
	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		temp := make([]Node, 0)
		for _, character := range line {
			temp = append(temp, Node{character, nil, nil, nil, nil, nil, nil, nil, nil})
		}
		nodes = append(nodes, temp)

		lineIndex++
	}

	for idy, row := range nodes {
		for idx, _ := range row {
			TY := idy - 1
			MY := idy
			BY := idy + 1
			LX := idx - 1
			MX := idx
			RX := idx + 1
			if TY > 0 {
				if LX > 0 {
					nodes[idy][idx].TL = &nodes[TY][LX]
				}
				nodes[idy][idx].TM = &nodes[TY][MX]
				if RX < len(row) {
					nodes[idy][idx].TR = &nodes[TY][RX]
				}
			}
			if LX > 0 {
				nodes[idy][idx].ML = &nodes[MY][LX]
			}
			if RX < len(row) {
				nodes[idy][idx].MR = &nodes[MY][RX]
			}
			if BY < len(nodes) {
				if LX > 0 {
					nodes[idy][idx].BL = &nodes[BY][LX]
				}
				nodes[idy][idx].BM = &nodes[BY][MX]
				if RX < len(row) {
					nodes[idy][idx].BR = &nodes[BY][RX]
				}
			}
		}
	}
	for _, row := range nodes {
		for _, node := range row {
			node.howManyXmas()
			return
		}
	}
}

// func partTwo(scanner bufio.Scanner) {
// }
