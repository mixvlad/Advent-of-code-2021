package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isLowest(i, j int, heights [][]int) bool {
	if i > 0 {
		if heights[i-1][j] <= heights[i][j] {
			return false
		}
	}
	if j > 0 {
		if heights[i][j-1] <= heights[i][j] {
			return false
		}
	}
	if i < len(heights)-1 {
		if heights[i+1][j] <= heights[i][j] {
			return false
		}
	}
	if j < len(heights[i])-1 {
		if heights[i][j+1] <= heights[i][j] {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	heights := make([][]int, 100)
	for i := range heights {
		heights[i] = make([]int, 100)
	}
	var row int
	for scanner.Scan() {

		curRow := scanner.Text()
		if len(curRow) > 0 {
			for k, v := range curRow {
				heights[row][k], _ = strconv.Atoi(string(v))
			}
		}
		row++
	}

	height := 0
	for i := range heights {
		for j := range heights[i] {
			if isLowest(i, j, heights) {
				height += heights[i][j] + 1
			}
		}
	}

	fmt.Println(height)
	secondPart()
}

func secondPart() {

}
