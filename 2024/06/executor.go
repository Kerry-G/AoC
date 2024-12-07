package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

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

type Node struct {
	kind     rune
	position [2]int
}

func partOne(scanner bufio.Scanner) {
	idy := 0
	visited := 0
	gameStruct := [][]Node{}
	direction := []string{"UP", "RIGHT", "DOWN", "LEFT"}
	directionState := 0
	var player *Node
	for scanner.Scan() {
		line := scanner.Text()
		gameStruct = append(gameStruct, []Node{})
		for idx, kind := range line {
			Node := Node{kind, [2]int{idx, idy}}
			gameStruct[idy] = append(gameStruct[idy], Node)
			if kind == '^' {
				player = &Node
			}

		}
		idy++
	}
	visited = 0
	run, seekPosition := true, [2]int{}
	for run {
		duration := time.Duration(0) * time.Millisecond
		time.Sleep(duration)
		run, seekPosition = traverse(player, &gameStruct, direction, &directionState)
		// fmt.Println()
		// for _, y := range gameStruct {
		// 	for _, x := range y {
		// 		fmt.Printf("%c", x.kind)
		// 	}
		// 	fmt.Println()
		// }

		seekNode := (gameStruct[seekPosition[1]][seekPosition[0]])
		gameStruct[player.position[1]][player.position[0]].kind = 'X'
		player = &seekNode
		gameStruct[player.position[1]][player.position[0]].kind = '^'

	}

	for _, y := range gameStruct {
		for _, x := range y {
			if x.kind == 'X' {
				visited++
			}
		}
	}

	fmt.Println(visited)
}

func traverse(player *Node, gameState *[][]Node, directionList []string, directionState *int) (bool, [2]int) {
	playerPosition := player.position
	var seekPosition [2]int
	isBlock := true
	for isBlock {
		direction := directionList[*directionState]
		if direction == "UP" {
			up := playerPosition[1] - 1
			seekPosition = [2]int{playerPosition[0], up}
		} else if direction == "RIGHT" {
			right := playerPosition[0] + 1
			seekPosition = [2]int{right, playerPosition[1]}
		} else if direction == "DOWN" {
			down := playerPosition[1] + 1
			seekPosition = [2]int{playerPosition[0], down}
		} else if direction == "LEFT" {
			left := playerPosition[0] - 1
			seekPosition = [2]int{left, playerPosition[1]}
		}
		boundY, boundX := len(*gameState), len((*gameState)[0])
		if seekPosition[0] < 0 || seekPosition[0] >= boundX || seekPosition[1] < 0 || seekPosition[1] >= boundY {
			// went out of bound
			(*gameState)[player.position[1]][player.position[0]].kind = 'X'
			return false, [2]int{}
		}
		if (*gameState)[seekPosition[1]][seekPosition[0]].kind == '#' {
			isBlock = true
			*directionState = (*directionState + 1) % (len(direction))
			fmt.Println(*directionState)
		} else {
			isBlock = false
		}
	}

	return true, seekPosition
}
