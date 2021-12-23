package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type point struct {
	i, j, cost, costFromStart int
	isVisited                 bool
}

type ByCostFromStart []point

func (a ByCostFromStart) Len() int           { return len(a) }
func (a ByCostFromStart) Less(i, j int) bool { return a[i].costFromStart < a[j].costFromStart }
func (a ByCostFromStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	fmt.Println("Part one:", calc(100, 1))
	fmt.Println("Part one:", calc(100, 5))
}

func Check(board [][]point, queue []point, x point, i, j int) []point {
	if i >= 0 && j >= 0 && i < len(board) && j < len(board[i]) {
		newCost := x.costFromStart + board[i][j].cost
		if board[i][j].costFromStart > newCost {
			board[i][j].costFromStart = newCost
		}
		if !board[i][j].isVisited {
			return append(queue, board[i][j])
		}
	}
	return queue
}

func calc(size, multiplier int) int {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	board := make([][]point, size*multiplier)
	for i := range board {
		board[i] = make([]point, size*multiplier)
	}

	var row int
	for scanner.Scan() {

		curRow := scanner.Text()
		if length := len(curRow); length > 0 {
			for k, v := range curRow {
				value, _ := strconv.Atoi(string(v))
				for i := 0; i < multiplier; i++ {
					xValue := (value + i)
					for {
						if xValue > 9 {
							xValue -= 9
						} else {
							break
						}
					}
					board[row][k+(i*length)] = point{row, k, xValue, MaxInt, false}
				}
			}
		}
		row++
	}
	for i := size; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			xValue := (board[i-size][j].cost + 1)
			if xValue > 9 {
				xValue = 1
			}
			board[i][j] = point{i, j, xValue, MaxInt, false}
		}
	}

	queue := []point{}

	board[0][0].costFromStart = 0
	board[0][0].cost = 0

	queue = append(queue, board[0][0])

	for {
		if len(queue) == 0 {
			break
		}

		sort.Sort(ByCostFromStart(queue))
		x := queue[0]
		if !board[x.i][x.j].isVisited {
			board[x.i][x.j].isVisited = true
			// check left
			queue = Check(board, queue, x, x.i-1, x.j)
			// check up
			queue = Check(board, queue, x, x.i, x.j+1)
			// check right
			queue = Check(board, queue, x, x.i+1, x.j)
			// check down
			queue = Check(board, queue, x, x.i, x.j-1)
		}

		queue = queue[1:]
	}

	return board[len(board)-1][len(board)-1].costFromStart
}
