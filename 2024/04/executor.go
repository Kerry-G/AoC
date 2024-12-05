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

	file, e := os.Open("resources/input.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	partTwo(*scanner)
}

type Node struct {
	element  rune
	TL       *Node
	TM       *Node
	TR       *Node
	ML       *Node
	MR       *Node
	BL       *Node
	BM       *Node
	BR       *Node
	position [2]int
}

func (node Node) howManyXmas() int {

	directions := []string{"TL", "TM", "TR", "ML", "MR", "BL", "BM", "BR"}
	// directions := []string{"TL"}
	acc := 0
	for _, direction := range directions {
		word, e := node.chain4(direction)
		if e != nil {

			continue
		} else {
			if word == "XMAS" {
				acc++
			}
		}
	}
	return acc
}

func (node Node) chain4(direction string) (string, error) {
	acc := ""
	for len(acc) != 4 {
		v := reflect.ValueOf(node)
		if len(acc) == 0 {
			acc = string(node.element)
		} else {
			acc = acc + string(node.element)
		}

		temp := v.FieldByName(direction)
		if temp.IsNil() {
			if len(acc) == 4 {
				continue
			}
			return "", fmt.Errorf("nil error")
		}
		nextNode := temp.Interface().(*Node)
		node = *nextNode

	}
	return acc, nil
}

func partOne(scanner bufio.Scanner) {
	nodes := make([][]Node, 0)
	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		temp := make([]Node, 0)
		for _, character := range line {
			temp = append(temp, Node{character, nil, nil, nil, nil, nil, nil, nil, nil, [2]int{}})
		}
		nodes = append(nodes, temp)

		lineIndex++
	}

	for idy, row := range nodes {
		for idx, _ := range row {
			nodes[idy][idx].position[0] = idy
			nodes[idy][idx].position[1] = idx
			TY := idy - 1
			MY := idy
			BY := idy + 1
			LX := idx - 1
			MX := idx
			RX := idx + 1
			if TY >= 0 {
				if LX >= 0 {
					nodes[idy][idx].TL = &nodes[TY][LX]
				}
				nodes[idy][idx].TM = &nodes[TY][MX]
				if RX < len(row) {
					nodes[idy][idx].TR = &nodes[TY][RX]
				}
			}
			if LX >= 0 {
				nodes[idy][idx].ML = &nodes[MY][LX]
			}
			if RX < len(row) {
				nodes[idy][idx].MR = &nodes[MY][RX]
			}
			if BY < len(nodes) {
				if LX >= 0 {
					nodes[idy][idx].BL = &nodes[BY][LX]
				}
				nodes[idy][idx].BM = &nodes[BY][MX]
				if RX < len(row) {
					nodes[idy][idx].BR = &nodes[BY][RX]
				}
			}
		}
	}
	sum := 0
	for _, row := range nodes {
		for _, node := range row {
			x := 'X'
			if node.element != x {
				continue
			}
			sum += node.howManyXmas()

		}
	}
	fmt.Println(sum)
}

func partTwo(scanner bufio.Scanner) {
	nodes := make([][]Node, 0)
	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		temp := make([]Node, 0)
		for _, character := range line {
			temp = append(temp, Node{character, nil, nil, nil, nil, nil, nil, nil, nil, [2]int{}})
		}
		nodes = append(nodes, temp)

		lineIndex++
	}

	for idy, row := range nodes {
		for idx, _ := range row {
			nodes[idy][idx].position[0] = idy
			nodes[idy][idx].position[1] = idx
			TY := idy - 1
			MY := idy
			BY := idy + 1
			LX := idx - 1
			MX := idx
			RX := idx + 1
			if TY >= 0 {
				if LX >= 0 {
					nodes[idy][idx].TL = &nodes[TY][LX]
				}
				nodes[idy][idx].TM = &nodes[TY][MX]
				if RX < len(row) {
					nodes[idy][idx].TR = &nodes[TY][RX]
				}
			}
			if LX >= 0 {
				nodes[idy][idx].ML = &nodes[MY][LX]
			}
			if RX < len(row) {
				nodes[idy][idx].MR = &nodes[MY][RX]
			}
			if BY < len(nodes) {
				if LX >= 0 {
					nodes[idy][idx].BL = &nodes[BY][LX]
				}
				nodes[idy][idx].BM = &nodes[BY][MX]
				if RX < len(row) {
					nodes[idy][idx].BR = &nodes[BY][RX]
				}
			}
		}
	}
	MAS_Nodes := [][]Node{}
	for _, row := range nodes {
		for _, node := range row {
			m := 'M'
			if node.element != m {
				continue
			}

			node.howManyMasCross(&MAS_Nodes)
		}
	}
	acc := 0
	a := 'A'

	for _, nodes := range MAS_Nodes {
		for _, compareNodes := range MAS_Nodes {
			if nodes[0].isEqual(compareNodes[0]) {
				continue
			}

			if nodes[1].element == a && (nodes[1].isEqual(compareNodes[1])) {
				acc++
			}
		}
	}
	// Ugly ass code, I don't removed counted MAS, therefore I got doubled of them (MAS and MAS are counted)
	fmt.Println(acc / 2)
}

func (node Node) howManyMasCross(MAS_Nodes *[][]Node) {

	directions := []string{"TL", "TR", "BL", "BR"}
	for _, direction := range directions {
		word, node_sequence, e := node.chain3(direction)
		if e != nil {
			continue
		} else {
			if word == "MAS" {
				*MAS_Nodes = append(*MAS_Nodes, node_sequence)
			}
		}
	}

}
func (node Node) isEqual(node2 Node) bool {
	return (node.position[0] == node2.position[0] && node.position[1] == node2.position[1])
}

func (node Node) chain3(direction string) (string, []Node, error) {
	acc := ""
	node_sequence := []Node{}

	for len(acc) != 3 {
		node_sequence = append(node_sequence, node)
		v := reflect.ValueOf(node)
		if len(acc) == 0 {
			acc = string(node.element)
		} else {
			acc = acc + string(node.element)
		}

		temp := v.FieldByName(direction)
		if temp.IsNil() {
			if len(acc) == 3 {
				continue
			}
			return "", nil, fmt.Errorf("nil error")
		}
		nextNode := temp.Interface().(*Node)
		node = *nextNode

	}
	return acc, node_sequence, nil
}
