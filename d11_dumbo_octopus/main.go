package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	board := make([][]int, 10)
	for i := range board {
		board[i] = make([]int, 10)
	}

	var row int
	for scanner.Scan() {

		curRow := scanner.Text()
		if len(curRow) > 0 {
			for k, v := range curRow {
				board[row][k], _ = strconv.Atoi(string(v))
			}
		}
		row++
	}

	var countFlashes int
	for i := 0; i < 1000; i++ {
		Step(board, &countFlashes)

		if i+1 == 100 {
			fmt.Println("Flashes after 100 steps:", countFlashes)
		}
		if chechAllFlash(board) {
			fmt.Println("All flash step:", i+1)
			break
		}
	}

}

func chechAllFlash(board [][]int) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] > 0 {
				return false
			}
		}
	}
	return true
}

func Step(board [][]int, flashes *int) {
	for i := range board {
		for j := range board[i] {
			board[i][j]++
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] > 9 {
				board[i][j] = 0
				*flashes++
				FlashAround(board, flashes, i, j)
			}
		}
	}
}

func FlashAround(board [][]int, flashes *int, i, j int) {
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x != i || y != j {
				CheckFlash(board, flashes, x, y)
			}
		}
	}
}

func CheckFlash(board [][]int, flashes *int, i, j int) {
	if i >= 0 && j >= 0 && i < len(board) && j < len(board[i]) {
		if board[i][j] > 0 {
			board[i][j]++
			if board[i][j] > 9 {
				board[i][j] = 0
				*flashes++
				FlashAround(board, flashes, i, j)
			}
		}
	}
}
